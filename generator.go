package goxsd

import (
	"bytes"
	"encoding/xml"
	"github.com/realmfoo/caementarii/xsd"
	"go/format"
	"io"
	"sort"
	"strings"
)

type Generator struct {
	PkgName        string
	ImportResolver func(namespace string, schemaLocation string) (*xsd.Schema, error)
	schemas        map[string]*schema
}

func (g *Generator) Generate(s *xsd.Schema, o io.Writer) error {
	schema, err := parseSchema(s, g)
	if err != nil {
		return err
	}

	file := toGoFile(g.PkgName, schema)
	w := new(bytes.Buffer)
	file.Write(w)

	formatted, err := format.Source(w.Bytes())
	if err != nil {
		o.Write(w.Bytes())
		return err
	}

	o.Write(formatted)
	return nil
}

type xmlNames []xml.Name

func (a xmlNames) Len() int           { return len(a) }
func (a xmlNames) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a xmlNames) Less(i, j int) bool { return a[i].Local < a[j].Local }

func toGoFile(pkgName string, schema *schema) *File {
	f := &File{PkgName: pkgName}

	// Sort elements by local name
	keys := make([]xml.Name, 0, len(schema.elementDeclarations))
	for k := range schema.elementDeclarations {
		keys = append(keys, k)
	}
	sort.Sort(xmlNames(keys))

	// Generate types in alphabetical order
	for _, key := range keys {
		elm := schema.elementDeclarations[key]
		typeName := makeTypeName(elm.name)
		decl := &TypeDecl{
			Name: &Name{Value: typeName},
			Type: createElementDeclType(f, elm, typeName),
		}

		f.DeclList = append(f.DeclList, decl)
	}

	return f
}

func makeTypeName(name xml.Name) string {
	return strings.Title(name.Local)
}

func createElementDeclType(f *File, elm *elementDeclaration, typeName string) Expr {
	var elmType Expr

	switch typeDef := elm.typeDefinition.(type) {
	case *simpleTypeDefinition:
		elmType = &Name{Value: typeDef.primitiveTypeDefinition.goType}

		// If it's a root type then add a namespace variable?
		if typeName != "" {
			f.Require("encoding/xml")
			f.DeclList = append(f.DeclList, &VarDecl{
				NameList: []*Name{{Value: "ns" + typeName}},
				Values:   &BasicLit{Value: `xml.Name{Space: "` + elm.name.Space + `", Local: "` + elm.name.Local + `"}`},
			})
		}
	case *complexTypeDefinition:
		elmType = createComplexTypeDeclType(f, elm, typeName, typeDef)
	}

	return elmType
}

func createComplexTypeDeclType(f *File, elm *elementDeclaration, typeName string, typeDef *complexTypeDefinition) *StructType {
	s := &StructType{
		FieldList: []*Field{},
	}

	// If it's a root type definition then add an XMLName field.
	if typeName != "" {
		f.Require("encoding/xml")
		s.FieldList = append(s.FieldList,
			&Field{
				Name: &Name{Value: "XMLName"},
				Type: &BasicLit{Value: `xml.Name`},
				Tags: map[string]string{
					"xml": xmlNameTag(elm.name),
				},
			},
		)
	}

	for _, attr := range typeDef.attributeUses {
		f.Require("encoding/xml")
		var attrType Expr
		tags := xmlNameTag(attr.attributeDeclaration.name) + ",attr"
		attrType = &BasicLit{Value: attr.attributeDeclaration.typeDefinition.goType}
		if !attr.required {
			tags += ",omitempty"
			attrType = &PointerType{Elem: attrType}
		}
		s.FieldList = append(s.FieldList,
			&Field{
				Name: &Name{Value: makeTypeName(attr.attributeDeclaration.name)},
				Type: attrType,
				Tags: map[string]string{
					"xml": tags,
				},
			},
		)
	}

	p := typeDef.contentType.particle
	if p != nil {
		switch term := p.term.(type) {
		case *modelGroup:
			// sequence
			if term.compositor == "sequence" {
				s.FieldList = append(s.FieldList, createSequenceDecls(term, f)...)
			}
		}
	}
	return s
}

func createSequenceDecls(term *modelGroup, f *File) []*Field {
	fields := make([]*Field, 0)
	for _, particle := range term.particles {
		switch tt := particle.term.(type) {
		case *elementDeclaration:
			dt := createElementDeclType(f, tt, "")
			if particle.maxOccurs > 1 {
				dt = &SliceType{Elem: dt}
			} else if particle.minOccurs == 0 {
				dt = &PointerType{Elem: dt}
			}
			fields = append(fields,
				&Field{
					Name: &Name{Value: makeTypeName(tt.name)},
					Type: dt,
					Tags: map[string]string{
						"xml": xmlNameTag(tt.name),
					},
				},
			)
		}
	}
	return fields
}

func xmlNameTag(name xml.Name) string {
	xn := ""
	if name.Space != "" {
		xn += name.Space + " "
	}
	xn += name.Local
	return xn
}
