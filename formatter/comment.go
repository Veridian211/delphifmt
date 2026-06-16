package formatter

import (
	"delphifmt/token"
	"fmt"
)

func (f *Formatter) writeComment(t *token.Token) {
	switch t.Type {
	case token.TokenLineComment:
		var start int
		for i := 2; i < len(t.Value); i++ {
			if t.Value[i] != ' ' {
				start = i
				break
			}
		}
		f.write("// " + t.Value[start:])
	case token.TokenBlockComment:
		switch t.Value[0] {
		case '{':
			var start int
			for i := 1; i < len(t.Value); i++ {
				if t.Value[i] != ' ' {
					start = i
					break
				}
			}
			var end int
			for i := len(t.Value) - 2; i > 1; i-- {
				if t.Value[i] != ' ' {
					end = i
					break
				}
			}
			f.write("{ " + t.Value[start:end+1] + " }")
		case '(':
			var start int
			for i := 2; i < len(t.Value); i++ {
				if t.Value[i] != ' ' {
					start = i
					break
				}
			}
			var end int
			for i := len(t.Value) - 3; i > 1; i-- {
				if t.Value[i] != ' ' {
					end = i
					break
				}
			}
			f.write("(* " + t.Value[start:end+1] + " *)")
		default:
			panic("Unreachable")
		}
	default:
		panic(fmt.Sprintf("Token is not a comment: %s", t.Value))
	}
}
