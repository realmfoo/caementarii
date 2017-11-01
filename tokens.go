package goxsd

import "fmt"

type token uint

const (
	_ token = iota
	_EOF

	// names and literals
	_Name
	_Literal

	// operators and operations
	_Operator // excluding '*' (_Star)
	_AssignOp
	_IncOp
	_Assign
	_Define
	_Arrow
	_Star

	// delimitors
	_Lparen
	_Lbrack
	_Lbrace
	_Rparen
	_Rbrack
	_Rbrace
	_Comma
	_Semi
	_Colon
	_Dot
	_DotDotDot

	// keywords
	_Break
	_Case
	_Chan
	_Const
	_Continue
	_Default
	_Defer
	_Else
	_Fallthrough
	_For
	_Func
	_Go
	_Goto
	_If
	_Import
	_Interface
	_Map
	_Package
	_Range
	_Return
	_Select
	_Struct
	_Switch
	_Type
	_Var

	tokenCount
)

var tokstrings = [...]string{
	// source control
	_EOF: "EOF",

	// names and literals
	_Name:    "name",
	_Literal: "literal",

	// operators and operations
	_Operator: "op",
	_AssignOp: "op=",
	_IncOp:    "opop",
	_Assign:   "=",
	_Define:   ":=",
	_Arrow:    "<-",
	_Star:     "*",

	// delimitors
	_Lparen:    "(",
	_Lbrack:    "[",
	_Lbrace:    "{",
	_Rparen:    ")",
	_Rbrack:    "]",
	_Rbrace:    "}",
	_Comma:     ",",
	_Semi:      ";",
	_Colon:     ":",
	_Dot:       ".",
	_DotDotDot: "...",

	// keywords
	_Break:       "break",
	_Case:        "case",
	_Chan:        "chan",
	_Const:       "const",
	_Continue:    "continue",
	_Default:     "default",
	_Defer:       "defer",
	_Else:        "else",
	_Fallthrough: "fallthrough",
	_For:         "for",
	_Func:        "func",
	_Go:          "go",
	_Goto:        "goto",
	_If:          "if",
	_Import:      "import",
	_Interface:   "interface",
	_Map:         "map",
	_Package:     "package",
	_Range:       "range",
	_Return:      "return",
	_Select:      "select",
	_Struct:      "struct",
	_Switch:      "switch",
	_Type:        "type",
	_Var:         "var",
}

func (tok token) String() string {
	var s string
	if 0 <= tok && int(tok) < len(tokstrings) {
		s = tokstrings[tok]
	}
	if s == "" {
		s = fmt.Sprintf("<tok-%d>", tok)
	}
	return s
}
