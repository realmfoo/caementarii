package goxsd

import (
	"bytes"
)

type LitKind uint

const (
	IntLit LitKind = iota
	FloatLit
	ImagLit
	RuneLit
	StringLit
)

// All declarations belonging to the same group point to the same Group node.
type Group struct {
	dummy int // not empty so we are guaranteed different Group instances
}

type Node interface {
	aNode()
}

type node struct{}

func (*node) aNode() {}

//-----------------------------------
// File
type (
	File struct {
		PkgName  string
		Imports  []Decl
		DeclList []Decl
	}
)

//-----------------------------------
// Declarations
type (
	ImportDecl struct {
		LocalPkgName *Name
		Path         *BasicLit
		Group        *Group // nil means not part of a group
		decl
	}

	// NameList Type
	// NameList Type = Values
	// NameList      = Values
	VarDecl struct {
		NameList []*Name
		Type     Expr   // nil means no type
		Values   Expr   // nil means no values
		Group    *Group // nil means not part of a group
		decl
	}

	TypeDecl struct {
		Name  *Name
		Alias bool
		Type  Expr
		Group *Group // nil means not part of a group
		decl
	}

	Decl interface {
		Node
		aDecl()
	}
)

type decl struct{ node }

func (*decl) aDecl() {}

//-----------------------------------
// Expressions

type (
	Expr interface {
		Node
		aExpr()
	}

	// Value
	Name struct {
		Value string
		expr
	}

	// Value
	BasicLit struct {
		Value string
		Kind  LitKind
		expr
	}

	// []Elem
	SliceType struct {
		Elem Expr
		expr
	}

	// *Elem
	PointerType struct {
		Elem Expr
		expr
	}

	// struct { FieldList[0] TagList[0]; FieldList[1] TagList[1]; ... }
	StructType struct {
		FieldList []*Field
		expr
	}

	// Name Type
	//      Type
	Field struct {
		Name *Name // nil means anonymous field/parameter (structs/parameters), or embedded interface (interfaces)
		Type Expr  // field names declared in a list share the same Type (identical pointers)
		Tags map[string]string
		node
	}

	// Type { ElemList[0], ElemList[1], ... }
	CompositeLit struct {
		Type     Expr // nil means no literal type
		ElemList []Expr
		NKeys    int // number of elements with keys
		expr
	}

	// Key: Value
	KeyValueExpr struct {
		Key, Value Expr
		expr
	}
)

type expr struct{ node }

func (*expr) aExpr() {}

//-----------------------------------
// Functions

func (f *File) Require(path string) {
	if !f.HasImport(path) {
		f.Imports = append(
			f.Imports,
			&ImportDecl{
				Group: &Group{},
				Path:  &BasicLit{Value: `"` + path + `"`},
			},
		)
	}
}

func (f *File) HasImport(path string) bool {
	for _, i := range f.Imports {
		if i.(*ImportDecl).Path.Value == `"`+path+`"` {
			return true
		}
	}

	return false
}

func (f *File) Write(buf *bytes.Buffer) {
	p := printer{output: buf}

	writePackageName(buf, f)
	p.print(&printGroup{Tok: _Import, Decls: f.Imports}, newline, newline)
	for _, t := range f.DeclList {
		p.print(t, newline)
	}

	p.flush(_EOF)
}

func writePackageName(w *bytes.Buffer, file *File) (int, error) {
	return w.WriteString("package " + file.PkgName + "\n\n")
}
