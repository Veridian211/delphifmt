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
	f.writeKeyword(node.ProgramKeyword)
	f.writeLn(
		" ",
		node.Name.Value,
		node.Semicolon.Value,
	)
	f.writeKeyword(node.Begin)
	f.writeLn()
	f.depth++
	f.formatStatements(node.Statements)
	f.depth--
	f.writeKeyword(node.End)
	f.writeLn(node.Dot.Value)
}

func (f *Formatter) formatStatements(nodes []*ast.StatementNode) {
	for _, node := range nodes {
		f.writeIndent()
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
	f.write(node.Name.Value)
	f.formatArgumentList(node.ArgumentList)
	f.write(node.Semicolon.Value)
}

func (f *Formatter) formatArgumentList(node ast.ArgumentListNode) {
	f.write(node.LeftParen.Value)
	for _, arg := range node.Args {
		f.formatArgument(*arg)
	}
	f.write(node.RightParen.Value)
}

func (f *Formatter) formatArgument(node ast.ArgNode) {
	f.write(node.Expression.Value)
	if node.Comma != nil {
		f.write(
			node.Comma.Value,
			" ",
		)
	}
}
