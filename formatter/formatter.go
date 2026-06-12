package formatter

import (
	"delphifmt/ast"
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
	f.writeLn(
		node.ProgramKeyword.Value,
		" ",
		node.Name.Value,
		node.Semicolon.Value,
	)
	f.writeLn(node.Begin.Value)
	f.depth++
	f.formatStatements(node.Statements)
	f.depth--
	f.writeLn(
		node.End.Value,
		node.Dot.Value,
	)
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
