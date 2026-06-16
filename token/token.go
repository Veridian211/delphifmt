package token

import "fmt"

type Token struct {
	Type           TokenType
	Value          string
	Line           int
	Col            int
	LeadingTrivia  []*Token
	TrailingTrivia []*Token
}

func CreateToken(
	t TokenType,
	value string,
	line int,
	col int,
) *Token {
	return &Token{
		Type:           t,
		Value:          value,
		Line:           line,
		Col:            col,
		LeadingTrivia:  make([]*Token, 0),
		TrailingTrivia: make([]*Token, 0),
	}
}

func (t *Token) IsKeyword() bool {
	return t.Type.IsKeyword()
}

func (t *Token) IsComment() bool {
	return t.Type == TokenLineComment ||
		t.Type == TokenBlockComment
}

func (t *Token) IsNewLine() bool {
	return t.Type == TokenNewline
}

func (t *Token) PrintDebugLn() {
	if len(t.LeadingTrivia) != 0 {
		for i, leading := range t.LeadingTrivia {
			if i == 0 {
				fmt.Printf("\t┌ %s\n", triviaToStr(leading))
				continue
			}
			fmt.Printf("\t├ %s\n", triviaToStr(leading))
		}
	}

	fmt.Printf(
		"Line: %d, Col: %d, Type: %s, Value: %s\n",
		t.Line,
		t.Col,
		t.Type.ToDebug(),
		t.Value,
	)

	if len(t.TrailingTrivia) != 0 {
		for i, trailing := range t.TrailingTrivia {
			if i == len(t.TrailingTrivia)-1 {
				fmt.Printf("\t└ %s\n", triviaToStr(trailing))
				continue
			}
			fmt.Printf("\t├ %s\n", triviaToStr(trailing))
		}
	}
}

func triviaToStr(t *Token) string {
	switch {
	case t.IsComment():
		return t.Value
	case t.IsNewLine():
		return "<newline>"
	default:
		panic("Is not a trivia token: ")
	}
}

func PrintDebugLn(tokens []Token) {
	for _, t := range tokens {
		t.PrintDebugLn()
	}
}
