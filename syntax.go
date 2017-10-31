package goxsd

import "bytes"

type (
	File struct {
		PkgName string
		Imports []ImportDecl
	}

	ImportDecl struct {
		LocalPkgName string
		Path         string
	}
)

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
		if imp.LocalPkgName != "" {
			w.WriteString(imp.LocalPkgName)
			w.WriteString(" ")
		}
		w.WriteString(`"` + imp.Path + `"` + "\n")
	}
	w.WriteString(")\n")
}
