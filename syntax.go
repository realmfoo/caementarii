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
		Imports  []ImportDecl
		DeclList []Decl
	}
)

//-----------------------------------
// Declarations
type (
	ImportDecl struct {
		LocalPkgName *Name
		Path         string
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
)

type expr struct{ node }

func (*expr) aExpr() {}

//-----------------------------------
// Functions

func (f *File) Require(path string) {
	if !f.HasImport(path) {
		f.Imports = append(f.Imports, ImportDecl{Path: "encoding/xml"})
	}
}

func (f *File) HasImport(path string) bool {
	for _, i := range f.Imports {
		if i.Path == path {
			return true
		}
	}

	return false
}

func (f *File) Write(buf *bytes.Buffer) {
	writePackageName(buf, f)
	writeImports(buf, f)
}

func writePackageName(w *bytes.Buffer, file *File) (int, error) {
	return w.WriteString("package " + file.PkgName + "\n\n")
}

func writeImports(w *bytes.Buffer, file *File) {
	w.WriteString("import (\n")
	for _, imp := range file.Imports {
		w.WriteString(`	`)
		if imp.LocalPkgName != nil {
			w.WriteString(imp.LocalPkgName.Value)
			w.WriteString(" ")
		}
		w.WriteString(`"` + imp.Path + `"` + "\n")
	}
	w.WriteString(")\n")
}
