package token

import "fmt"

type Token struct {
	Type             TokenType
	Value            string
	Line             int
	Col              int
	LeadingComments  []*Token
	TrailingComments []*Token
}

func CreateToken(
	t TokenType,
	value string,
	line int,
	col int,
) *Token {
	return &Token{
		Type:             t,
		Value:            value,
		Line:             line,
		Col:              col,
		LeadingComments:  make([]*Token, 0),
		TrailingComments: make([]*Token, 0),
	}
}

func (t *Token) IsKeyword() bool {
	return t.Type.IsKeyword()
}

func (t *Token) IsComment() bool {
	return t.Type == TokenLineComment ||
		t.Type == TokenBlockComment
}

func (t *Token) PrintDebugLn() {
	if len(t.LeadingComments) != 0 {
		for i, leading := range t.LeadingComments {
			if i == 0 {
				fmt.Printf("\t┌ %s\n", leading.Value)
				continue
			}
			fmt.Printf("\t├ %s\n", leading.Value)
		}
	}
	fmt.Printf(
		"Line: %d, Col: %d, Type: %s, Value: %s\n",
		t.Line,
		t.Col,
		t.Type.ToDebug(),
		t.Value,
	)
	if len(t.TrailingComments) != 0 {
		for i, trailing := range t.TrailingComments {
			if i == len(t.TrailingComments)-1 {
				fmt.Printf("\t└ %s\n", trailing.Value)
				continue
			}
			fmt.Printf("\t├ %s\n", trailing.Value)
		}
	}
}

func PrintDebugLn(tokens []Token) {
	for _, t := range tokens {
		t.PrintDebugLn()
	}
}
