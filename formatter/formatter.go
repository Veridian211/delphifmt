package formatter

import (
	"delphifmt/ast"
	"delphifmt/token"
	"fmt"
	"strings"
)

type Formatter struct {
	options Options
	depth   int
	output  strings.Builder
}

func NewFormatter() *Formatter {
	return &Formatter{
		output:  strings.Builder{},
		options: DefaultOptions(),
	}
}

func (f *Formatter) writeToken(t *token.Token) {
	for _, leading := range t.LeadingComments {
		f.writeIndent()
		f.writeComment(leading)
		f.writeLn()
	}

	if f.output.Len() > 0 {
		lastChar := f.output.String()[f.output.Len()-1]
		if lastChar == '\n' {
			f.writeIndent()
		}
	}

	if t.Type.IsKeyword() {
		f.writeKeyword(t)
	} else {
		f.write(t.Value)
	}

	for _, trailing := range t.TrailingComments {
		f.write(" ")
		f.writeComment(trailing)
		if trailing.Type == token.TokenLineComment {
			f.writeLn()
		}
	}
}
func (f *Formatter) writeKeyword(t *token.Token) {
	var keyword string
	switch f.options.KeywordCase {
	case KeywordLowercase:
		keyword = strings.ToLower(t.Value)
	case KeywordUppercase:
		keyword = strings.ToUpper(string(t.Value[0])) + strings.ToLower(t.Value[1:])
	default:
		panic(fmt.Sprintf("Unknown option for keywordCase: %s", t.Value))
	}
	f.write(keyword)
}

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
			f.write(t.Value)
		default:
			panic("Unreachable")
		}
	default:
		panic(fmt.Sprintf("Token is not a comment: %s", t.Value))
	}
}

func (f *Formatter) writeIndent() {
	for range f.depth {
		switch f.options.IndentChar {
		case IndentSpace:
			for range f.options.Indent {
				f.write(" ")
			}
		case IndentTab:
			f.write("\t")
		}
	}
}

func (f *Formatter) write(input ...string) {
	for _, s := range input {
		f.output.WriteString(s)
	}
}

func (f *Formatter) writeLn(input ...string) {
	for _, s := range input {
		f.output.WriteString(s)
	}
	f.output.WriteString("\n")
}

func (f *Formatter) Format(node ast.Node) string {
	f.output.Reset()
	f.formatInternal(node)
	return f.output.String()
}

func (f *Formatter) formatInternal(node ast.Node) {
	switch n := node.(type) {
	case *ast.ProgramNode:
		f.formatProgramNode(*n)
	}
}

func (f *Formatter) formatProgramNode(node ast.ProgramNode) {
	f.writeToken(node.ProgramKeyword)
	f.write(" ")
	f.writeToken(node.Name)
	f.writeToken(node.Semicolon)
	f.writeLn()

	f.writeToken(node.Begin)
	f.writeLn()
	f.depth++
	f.formatStatements(node.Statements)
	f.depth--
	f.writeToken(node.End)
	f.writeToken(node.Dot)
	f.writeLn()
}

func (f *Formatter) formatStatements(nodes []*ast.StatementNode) {
	for _, node := range nodes {
		f.formatStatement(*node)
		f.writeLn()
	}
}

func (f *Formatter) formatStatement(node ast.StatementNode) {
	switch n := node.(type) {
	case *ast.MethodCallNode:
		f.formatMethodCall(*n)
	}
}

func (f *Formatter) formatMethodCall(node ast.MethodCallNode) {
	f.writeToken(node.Name)
	f.formatArgumentList(node.ArgumentList)
	f.writeToken(node.Semicolon)
}

func (f *Formatter) formatArgumentList(node ast.ArgumentListNode) {
	f.writeToken(node.LeftParen)
	for _, arg := range node.Args {
		f.formatArgument(*arg)
	}
	f.writeToken(node.RightParen)
}

func (f *Formatter) formatArgument(node ast.ArgNode) {
	f.writeToken(node.Expression)
	if node.Comma != nil {
		f.writeToken(node.Comma)
		f.write(" ")
	}
}
