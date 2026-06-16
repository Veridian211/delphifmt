package ast

import "delphifmt/token"

type NodeKind int

const (
	NodeProgram = iota

	NodeStatement
	NodeAssignment
	NodeMethodCall
	NodeArgumentList
	NodeArgument
	NodeIfStatement
	NodeForLoop
	NodeWhileLoop
	NodeDoWhileLoop
)

type Node interface {
	NodeKind() NodeKind
}

// program

type ProgramNode struct {
	ProgramKeyword *token.Token
	Name           *token.Token
	Semicolon      *token.Token
	Begin          *token.Token
	Statements     []*StatementNode
	End            *token.Token
	Dot            *token.Token
	Eof            *token.Token
}

func (*ProgramNode) NodeKind() NodeKind { return NodeProgram }

// statements

type StatementNode interface {
	Node
	statementNode()
}

type MethodCallNode struct {
	Name         *token.Token
	ArgumentList ArgumentListNode
	Semicolon    *token.Token
}

func (*MethodCallNode) NodeKind() NodeKind { return NodeMethodCall }
func (*MethodCallNode) statementNode()     {}

type ArgumentListNode struct {
	LeftParen  *token.Token
	Args       []*ArgNode
	RightParen *token.Token
}

func (*ArgumentListNode) NodeKind() NodeKind { return NodeArgumentList }
func (*ArgumentListNode) statementNode()     {}

type ArgNode struct {
	Expression *token.Token
	Comma      *token.Token // optional
}

func (*ArgNode) NodeKind() NodeKind { return NodeArgument }
func (*ArgNode) statementNode()     {}
