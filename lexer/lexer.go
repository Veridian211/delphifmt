package lexer

import (
	"delphifmt/token"
	"fmt"
	"strings"
)

type Lexer struct {
	src  string
	pos  int
	line int
	col  int
}

func NewLexer(src string) *Lexer {
	return &Lexer{src: src, line: 1, col: 1}
}

func (l *Lexer) LexSrc() []token.Token {
	tokens := make([]token.Token, 0, 1024)
	for {
		tok := l.nextToken()
		tokens = append(tokens, tok)

		if tok.Type == token.TokenEOF {
			break
		}
	}
	tokens = l.appendTrivia(tokens)
	return tokens
}

func (l *Lexer) peek() byte {
	if l.pos >= len(l.src) {
		return 0
	}
	return l.src[l.pos]
}

func (l *Lexer) peekNext() byte {
	if l.pos+1 >= len(l.src) {
		return 0
	}
	return l.src[l.pos+1]
}

func (l *Lexer) advance() byte {
	b := l.src[l.pos]
	l.pos++
	if b == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}
	return b
}

func (l *Lexer) tok(t token.TokenType, value string) token.Token {
	return *token.CreateToken(
		t,
		value,
		l.line,
		l.col-len(value),
	)
}

func (l *Lexer) nextToken() token.Token {
	if l.pos >= len(l.src) {
		return l.tok(token.TokenEOF, "")
	}

	ch := l.peek()

	switch {
	case isWhitespace(ch):
		if ch == '\n' && l.peekNext() == '\n' {
			l.advance()
			return l.tok(token.TokenNewline, string(ch))
		}
		l.advance()
		return l.nextToken()

	case ch == ';':
		l.advance()
		return l.tok(token.TokenSemicolon, string(ch))

	case ch == '.':
		l.advance()
		return l.tok(token.TokenDot, string(ch))

	case ch == '(':
		if l.peekNext() == '*' {
			return l.readParenBlockComment()
		}
		l.advance()
		return l.tok(token.TokenParenLeft, string(ch))

	case ch == ')':
		l.advance()
		return l.tok(token.TokenParenRight, string(ch))

	case ch == '/':
		if l.peekNext() == '/' {
			return l.readLineComment()
		}
		l.advance()
		return l.tok(token.TokenSlash, string(ch))

	case ch == '{':
		return l.readCurlyBlockComment()

	case ch == '\'':
		return l.readString()

	case isLetter(ch):
		return l.readIdentifier()

	case isDigit(ch):
		return l.readNumber()

	default:
		panic(fmt.Errorf("Unexpected char \"%c\" at line %d, col %d", ch, l.line, l.col))
	}
}

func (l *Lexer) readIdentifier() token.Token {
	start := l.pos
	for l.pos < len(l.src) && (isLetter(l.src[l.pos]) || isDigit(l.src[l.pos])) {
		l.advance()
	}
	ident := l.src[start:l.pos]
	switch {
	case strings.EqualFold(ident, "program"):
		return l.tok(token.TokenProgram, ident)
	case strings.EqualFold(ident, "begin"):
		return l.tok(token.TokenBegin, ident)
	case strings.EqualFold(ident, "end"):
		return l.tok(token.TokenEnd, ident)
	}
	return l.tok(token.TokenIdentifier, ident)
}

func (l *Lexer) readNumber() token.Token {
	start := l.pos
	for l.pos < len(l.src) && isDigit(l.src[l.pos]) {
		l.advance()
	}
	return l.tok(token.TokenNumber, l.src[start:l.pos])
}

func (l *Lexer) readString() token.Token {
	start := l.pos
	l.advance()
	for l.pos < len(l.src) {
		ch := l.advance()
		if ch == '\'' {
			if l.peek() == '\'' {
				l.advance()
				continue
			}
			break
		}
	}
	return l.tok(token.TokenString, l.src[start:l.pos])
}

func (l *Lexer) readLineComment() token.Token {
	start := l.pos
	l.advance()
	for l.pos < len(l.src) && l.src[l.pos] != '\n' {
		l.advance()
	}
	return l.tok(token.TokenLineComment, l.src[start:l.pos])
}

func (l *Lexer) readParenBlockComment() token.Token {
	start := l.pos
	l.advance() // skip (
	l.advance() // skip *
	for l.pos < len(l.src) {
		ch := l.advance()
		if ch == '*' {
			if l.peek() == ')' {
				l.advance()
				break
			}
		}
	}
	return l.tok(token.TokenBlockComment, l.src[start:l.pos])
}

func (l *Lexer) readCurlyBlockComment() token.Token {
	start := l.pos
	l.advance()
	for l.pos < len(l.src) && l.src[l.pos] != '}' {
		l.advance()
	}
	l.advance()
	return l.tok(token.TokenBlockComment, l.src[start:l.pos])
}

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func (*Lexer) appendTrivia(tokens []token.Token) []token.Token {
	result := make([]token.Token, 0, len(tokens))
	i := 0
	for i < len(tokens) {
		var leadingTrivia []*token.Token
		for i < len(tokens) && (tokens[i].IsComment() || tokens[i].IsNewLine()) {
			leadingTrivia = append(leadingTrivia, &tokens[i])
			i++
		}
		if i >= len(tokens) {
			break
		}

		tok := tokens[i]
		tok.LeadingTrivia = leadingTrivia
		i++

		for i < len(tokens) && tokens[i].IsComment() {
			if tokens[i].Line > tok.Line {
				break
			}
			tok.TrailingTrivia = append(tok.TrailingTrivia, &tokens[i])
			i++
		}
		result = append(result, tok)
	}
	return result
}
