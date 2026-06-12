package token

type TokenType int

const (
	TokenEOF = iota

	TokenSemicolon
	TokenDot

	TokenProgram
	TokenBegin
	TokenEnd

	TokenParenLeft
	TokenParenRight

	TokenIdentifier
	TokenString
	TokenNumber
)

var TokenTypeStr = map[TokenType]string{
	TokenEOF: "<EOF>",

	TokenSemicolon: ";",
	TokenDot:       ".",

	TokenProgram: "program",
	TokenBegin:   "begin",
	TokenEnd:     "end",

	TokenParenLeft:  "(",
	TokenParenRight: ")",

	TokenIdentifier: "<ident>",
	TokenString:     "<string>",
	TokenNumber:     "<number>",
}

func (t *TokenType) ToDebug() string {
	return TokenTypeStr[*t]
}
