package token

import "fmt"

type Token struct {
	Type  TokenType
	Value string
	Line  int
	Col   int
}

func (t *Token) PrintDebugLn() {
	fmt.Printf(
		"Line: %d  Col: %d  Type: %s  Value: %s\n",
		t.Line,
		t.Col,
		t.Type.ToDebug(),
		t.Value,
	)
}
