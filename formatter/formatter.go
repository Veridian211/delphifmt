package formatter

import (
	"delphifmt/ast"
	"delphifmt/token"
	"fmt"
	"strings"
)

type Formatter struct {
	Options Options
	depth   int
	output  strings.Builder
}

func NewFormatter() *Formatter {
	return &Formatter{
		output:  strings.Builder{},
		Options: DefaultOptions(),
	}
}

func (f *Formatter) writeToken(t *token.Token) {
	for _, leading := range t.LeadingTrivia {
		if leading.IsNewLine() {
			f.writeLn()
			continue
		}
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

	for _, trailing := range t.TrailingTrivia {
		f.write(" ")
		f.writeComment(trailing)
		if trailing.Type == token.TokenLineComment {
			f.writeLn()
		}
	}
}
func (f *Formatter) writeKeyword(t *token.Token) {
	var keyword string
	switch f.Options.KeywordCase {
	case KeywordLowercase:
		keyword = strings.ToLower(t.Value)
	case KeywordUppercase:
		keyword = strings.ToUpper(string(t.Value[0])) + strings.ToLower(t.Value[1:])
	default:
		panic(fmt.Sprintf("Unknown option for keywordCase: %s", t.Value))
	}
	f.write(keyword)
}

func (f *Formatter) writeIndent() {
	for range f.depth {
		switch f.Options.IndentChar {
		case IndentSpace:
			for range f.Options.Indent {
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

func (f *Formatter) lastChar() byte {
	if f.output.Len() == 0 {
		return 0
	}
	return f.output.String()[f.output.Len()-1]
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
	if f.lastChar() != '\n' {
		f.write(" ")
	}
	f.writeToken(node.Name)
	f.writeToken(node.Semicolon)
	if f.lastChar() != '\n' {
		f.writeLn()
	}

	f.writeToken(node.Begin)
	if f.lastChar() != '\n' {
		f.writeLn()
	}
	f.depth++
	f.formatStatements(node.Statements)
	f.depth--
	f.writeToken(node.End)
	f.writeToken(node.Dot)
	if f.lastChar() != '\n' {
		f.writeLn()
	}
	f.writeToken(node.Eof)
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
	f.depth++
	for _, arg := range node.Args {
		f.formatArgument(*arg)
	}
	f.depth--
	f.writeToken(node.RightParen)
}

func (f *Formatter) formatArgument(node ast.ArgNode) {
	f.writeToken(node.Expression)
	if node.Comma != nil {
		f.writeToken(node.Comma)
		f.write(" ")
	}
}
