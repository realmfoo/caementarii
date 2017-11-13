package goxsd

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/realmfoo/caementarii/xsd"
	"go/format"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Generator struct {
	PkgName string
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

func toGoFile(pkgName string, schema *schema) *File {
	f := &File{PkgName: pkgName}

	for _, elm := range schema.elementDeclarations {
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

		p := typeDef.contentType.particle
		if p != nil {
			switch term := p.term.(type) {
			case *modelGroup:
				// sequence
				if term.compositor == "sequence" {
					for _, particle := range term.particles {
						switch tt := particle.term.(type) {
						case *elementDeclaration:
							dt := createElementDeclType(f, tt, "")
							if particle.maxOccurs > 1 {
								dt = &SliceType{Elem: dt}
							} else if particle.minOccurs == 0 {
								dt = &PointerType{Elem: dt}
							}
							s.FieldList = append(s.FieldList,
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
				}
			}
		}

		elmType = s
	}

	return elmType
}

func xmlNameTag(name xml.Name) string {
	xn := ""
	if name.Space != "" {
		xn += name.Space + " "
	}
	xn += name.Local
	return xn
}
