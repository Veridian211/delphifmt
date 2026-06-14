package token

import "fmt"

type TokenType int

const (
	TokenEOF = iota

	TokenLineComment
	TokenBlockComment

	TokenSemicolon
	TokenDot
	TokenComma
	TokenColon
	TokenCaret
	TokenAt
	TokenAmpersand
	TokenEquals
	TokenAssign

	TokenProgram
	TokenBegin
	TokenEnd

	TokenUnit
	TokenInterface
	TokenUses
	TokenImplementation
	TokenInitialization
	TokenFinalization

	TokenLibrary
	TokenExports

	TokenType_
	TokenArray
	TokenSet
	TokenFile
	TokenClass
	TokenRecord
	TokenPacked

	TokenConstructor
	TokenDestructor
	TokenInherited
	TokenProperty

	TokenProcedure
	TokenFunction

	TokenVar
	TokenConst
	TokenOut
	TokenThreadVar
	TokenResourcestring

	TokenIf
	TokenThen
	TokenElse

	TokenCase
	TokenOf

	TokenFor
	TokenTo
	TokenDownTo
	TokenIn
	TokenDo
	TokenWhile
	TokenRepeat
	TokenUntil

	TokenWith

	TokenGoTo
	TokenLabel

	TokenTry
	TokenFinally
	TokenExcept
	TokenRaise

	TokenParenLeft
	TokenParenRight
	TokenBracketLeft
	TokenBracketRight
	TokenAngleLeft
	TokenAngleRight

	TokenNot
	TokenAnd
	TokenOr
	TokenXOr

	TokenIs

	TokenPlus
	TokenMinus
	TokenAsterisk
	TokenSlash
	TokenMod
	TokenDiv
	TokenShl
	TokenShr

	TokenNil

	TokenIdentifier
	TokenString
	TokenNumber
)

var TokenTypeStr = map[TokenType]string{
	TokenEOF: "<EOF>",

	TokenLineComment:  "<line-comment>",
	TokenBlockComment: "<block-comment>",

	TokenSemicolon: ";",
	TokenDot:       ".",
	TokenComma:     ",",
	TokenColon:     ":",
	TokenCaret:     "^",
	TokenAt:        "@",
	TokenAmpersand: "&",
	TokenEquals:    "=",
	TokenAssign:    ":=",

	TokenProgram: "program",
	TokenBegin:   "begin",
	TokenEnd:     "end",

	TokenUnit:           "unit",
	TokenInterface:      "interface",
	TokenUses:           "uses",
	TokenImplementation: "implementation",
	TokenInitialization: "initialization",
	TokenFinalization:   "finalization",
	TokenLibrary:        "library",
	TokenExports:        "exports",

	TokenType_:  "type",
	TokenArray:  "array",
	TokenSet:    "set",
	TokenFile:   "file",
	TokenClass:  "class",
	TokenRecord: "record",
	TokenPacked: "packed",

	TokenConstructor: "constructor",
	TokenDestructor:  "destructor",
	TokenInherited:   "inherited",
	TokenProperty:    "property",

	TokenProcedure: "procedure",
	TokenFunction:  "function",

	TokenVar:            "var",
	TokenConst:          "const",
	TokenOut:            "out",
	TokenThreadVar:      "threadvar",
	TokenResourcestring: "resourcestring",

	TokenIf:   "if",
	TokenThen: "then",
	TokenElse: "else",

	TokenCase: "case",
	TokenOf:   "of",

	TokenFor:    "for",
	TokenTo:     "to",
	TokenDownTo: "downto",
	TokenIn:     "in",
	TokenDo:     "do",
	TokenWhile:  "while",
	TokenRepeat: "repeat",
	TokenUntil:  "until",

	TokenWith: "with",

	TokenGoTo:  "goto",
	TokenLabel: "label",

	TokenTry:     "try",
	TokenFinally: "finally",
	TokenExcept:  "except",
	TokenRaise:   "raise",

	TokenParenLeft:    "(",
	TokenParenRight:   ")",
	TokenBracketLeft:  "[",
	TokenBracketRight: "]",
	TokenAngleLeft:    "<",
	TokenAngleRight:   ">",

	TokenNot: "not",
	TokenAnd: "and",
	TokenOr:  "or",
	TokenXOr: "xor",

	TokenIs: "is",

	TokenPlus:     "+",
	TokenMinus:    "-",
	TokenAsterisk: "*",
	TokenSlash:    "/",
	TokenMod:      "mod",
	TokenDiv:      "div",
	TokenShl:      "shl",
	TokenShr:      "shr",

	TokenNil:        "nil",
	TokenIdentifier: "<ident>",
	TokenString:     "<string>",
	TokenNumber:     "<number>",
}

func (t TokenType) IsKeyword() bool {
	switch t {
	case TokenProgram, TokenBegin, TokenEnd,
		TokenUnit, TokenInterface, TokenUses, TokenImplementation, TokenInitialization, TokenFinalization,
		TokenLibrary, TokenExports,
		TokenType_, TokenArray, TokenSet, TokenFile, TokenClass, TokenRecord, TokenPacked,
		TokenConstructor, TokenDestructor, TokenInherited, TokenProperty,
		TokenProcedure, TokenFunction,
		TokenVar, TokenConst, TokenOut, TokenThreadVar, TokenResourcestring,
		TokenIf, TokenThen, TokenElse,
		TokenCase, TokenOf,
		TokenFor, TokenTo, TokenDownTo, TokenIn, TokenDo, TokenWhile, TokenRepeat, TokenUntil,
		TokenWith,
		TokenGoTo, TokenLabel,
		TokenTry, TokenFinally, TokenExcept, TokenRaise,
		TokenNot, TokenAnd, TokenOr, TokenXOr,
		TokenIs, TokenMod, TokenDiv, TokenShl, TokenShr,
		TokenNil:
		return true
	}
	return false
}

func (t TokenType) ToDebug() string {
	str, ok := TokenTypeStr[t]
	if !ok {
		panic(fmt.Sprintf("Unknown TokenType: %d", t))
	}
	return str
}
