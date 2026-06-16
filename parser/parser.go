package parser

import (
	"delphifmt/ast"
	"delphifmt/token"
	"fmt"
)

type Parser struct {
	tokens []token.Token
	pos    int
	errors []ParseError
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens: tokens,
		errors: make([]ParseError, 0, 1024),
	}
}

func (p *Parser) peek() *token.Token {
	return &p.tokens[p.pos]
}

func (p *Parser) consume() *token.Token {
	t := p.tokens[p.pos]
	p.pos++
	return &t
}

func (p *Parser) expect(t token.TokenType) *token.Token {
	token := p.consume()
	if token.Type != t {
		p.errors = append(p.errors, ParseError{
			line: token.Line,
			col:  token.Col,
			message: fmt.Sprintf(
				"Expected %s, got %s",
				t.ToDebug(),
				token.Type.ToDebug(),
			),
		})
		// TODO: maybe return error token
	}
	return token
}

func (p *Parser) printPeek() {
	p.peek().PrintDebugLn()
}

func (p *Parser) GetErrors() []ParseError {
	return p.errors
}

func (p *Parser) hasNoErrors() bool {
	return len(p.errors) == 0
}

func (p *Parser) ParseProgram() (ast.ProgramNode, bool) {
	node := ast.ProgramNode{}
	node.ProgramKeyword = p.expect(token.TokenProgram)
	node.Name = p.expect(token.TokenIdentifier)
	node.Semicolon = p.expect(token.TokenSemicolon)
	node.Begin = p.expect(token.TokenBegin)
	for p.peek().Type != token.TokenEnd && p.peek().Type != token.TokenEOF {
		stmt := p.ParseStatement()
		node.Statements = append(node.Statements, &stmt)
	}
	node.End = p.expect(token.TokenEnd)
	node.Dot = p.expect(token.TokenDot)
	node.Eof = p.expect(token.TokenEOF)
	return node, p.hasNoErrors()
}

func (p *Parser) ParseStatement() ast.StatementNode {
	node := ast.MethodCallNode{}
	node.Name = p.expect(token.TokenIdentifier)
	node.ArgumentList = p.ParseArgumentList()
	node.Semicolon = p.expect(token.TokenSemicolon)
	return &node
}

func (p *Parser) ParseArgumentList() ast.ArgumentListNode {
	node := ast.ArgumentListNode{}
	node.LeftParen = p.expect(token.TokenParenLeft)
	for p.peek().Type != token.TokenEOF && p.peek().Type != token.TokenParenRight {
		arg := p.ParseArgument()
		node.Args = append(node.Args, &arg)
	}
	node.RightParen = p.expect(token.TokenParenRight)
	return node
}

func (p *Parser) ParseArgument() ast.ArgNode {
	node := ast.ArgNode{}
	node.Expression = p.expect(token.TokenString)
	return node
}
