package parser

import "fmt"

type Directive int

const (
	DirPackage = iota
	DirRequires
	DirContains

	DirForward

	DirAbstract
	DirSealed
	DirHelper

	DirPublic
	DirPrivate
	DirProtected
	DirStrict

	DirOverride
	DirReintroduce
	DirVirtual
	DirOverload

	DirRead
	DirWrite
	DirDefault
	DirNoDefault
	DirImplements

	DirOperator // like in class operator

	DirDeprecated
	DirPlatform
	DirExperimental
	DirLibrary

	DirStatic
)

var DirectiveStr = map[Directive]string{
	DirPackage:     "package",
	DirRequires:    "requires",
	DirContains:    "contains",

	DirForward:     "forward",

	DirAbstract:    "abstract",
	DirSealed:      "sealed",
	DirHelper:      "helper",

	DirPublic:      "public",
	DirPrivate:     "private",
	DirProtected:   "protected",
	DirStrict:      "strict",

	DirOverride:    "override",
	DirReintroduce: "reintroduce",
	DirVirtual:     "virtual",
	DirOverload:    "overload",

	DirRead:        "read",
	DirWrite:       "write",
	DirDefault:     "default",
	DirNoDefault:   "nodefault",
	DirImplements:  "implements",

	DirOperator:    "operator",

	DirDeprecated:  "deprecated",
	DirPlatform:    "platform",
	DirExperimental: "experimental",
	DirLibrary:     "library",

	DirStatic:      "static",
}

func (t Directive) ToDebug() string {
	str, ok := DirectiveStr[t]
	if !ok {
		panic(fmt.Sprintf("Unknown Directive: %d", t))
	}
	return str
}
