package goxsd

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/realmfoo/caementarii/xsd"
	"go/format"
	"io"
	"regexp"
	"strings"
)

type Generator struct {
	PkgName string
}

func (g *Generator) Generate(s *xsd.Schema, o io.Writer) error {
	schema := parseSchema(s, g)

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
		decl := &TypeDecl{
			Name: &Name{Value: strings.Title(elm.name.Local)},
		}

		if typeDef, ok := elm.typeDefinition.(*simpleTypeDefinition); ok {
			decl.Type = &Name{Value: typeDef.primitiveTypeDefinition.goType}
			f.Require("encoding/xml")
			f.DeclList = append(f.DeclList, &VarDecl{
				NameList: []*Name{{Value: "ns" + decl.Name.Value}},
				Values: &CompositeLit{
					Type: &Name{Value: "xml.Name"},
					ElemList: []Expr{
						&KeyValueExpr{
							Key:   &BasicLit{Value: "Space"},
							Value: &BasicLit{Value: `"` + elm.name.Space + `"`},
						},
						&KeyValueExpr{
							Key:   &BasicLit{Value: "Local"},
							Value: &BasicLit{Value: `"` + elm.name.Local + `"`},
						},
					},
				},
			})
		} else {
			decl.Type = &StructType{}
		}

		f.DeclList = append(f.DeclList, decl)
	}

	return f
}

func parseSchema(s *xsd.Schema, g *Generator) *schema {
	schema := newSchema(s)
	for _, top := range s.SchemaTop {
		if node, ok := top.(xsd.Element); ok {
			// 3.3.2.1 Common Mapping Rules for Element Declarations
			elm := g.newElement(schema, node)

			// 3.3.2.2 Mapping Rules for Top-Level Element Declarations
			elm.name.Space = s.TargetNamespace

			elm.scope.variety = "global"

			schema.elementDeclarations[elm.name] = elm
		}
	}
	return schema
}

func (g *Generator) newComplexType(s *schema, parent interface{}, node *xsd.ComplexType) *complexTypeDefinition {
	typeDef := complexTypeDefinition{}

	// The ·actual value· of the name [attribute].
	typeDef.name.Local = node.Name
	// The ·actual value· of the targetNamespace [attribute] of the <schema> ancestor element information item if present, otherwise ·absent·.
	typeDef.name.Space = s.targetNamespace

	s.typeDefinitions[typeDef.name] = &typeDef

	// The ·actual value· of the abstract [attribute], if present, otherwise false.
	typeDef.abstract = node.Abstract

	typeDef.prohibitedSubstitutions = getBlocks(node, s, typeDef)
	typeDef.final = getFinals(node, s, typeDef)

	// If the name [attribute] is present, then ·absent·, otherwise (among the ancestor element information items there
	// will be a nearest <element>), the Element Declaration corresponding to the nearest <element> information item
	// among the the ancestor element information items.
	if node.Name == "" {
		typeDef.context = parent
	}

	// A sequence whose members are Assertions drawn from the following sources, in order:
	// 1 The {assertions} of the {base type definition}.
	// 2 Assertions corresponding to all the <assert> element information items among the [children] of <complexType>, <restriction> and <extension>, if any, in document order.
	//typeDefinition.assertions =

	// The ·annotation mapping· of the set of elements containing the <complexType>, the <openContent> [child], if
	// present, the <attributeGroup> [children], if present, the <simpleContent> and <complexContent> [children], if
	// present, and their <restriction> and <extension> [children], if present, and their <openContent> and
	// <attributeGroup> [children], if present, as defined in
	// XML Representation of Annotation Schema Components (§3.15.2).
	//typeDefinition.annotations =

	if node.SimpleContent != nil {
		// If the <restriction> alternative is chosen, then restriction, otherwise (the <extension> alternative is
		// chosen) extension.
		if node.SimpleContent.Restriction != nil {
			// The type definition ·resolved· to by the ·actual value· of the base [attribute] on the <restriction> or
			// <extension> element appearing as a child of <simpleContent>
			typeDef.baseTypeDefinition = g.resolveType(s, s.resolveQName(node.SimpleContent.Restriction.Base))
			typeDef.derivationMethod = "restriction"
		} else {
			// The type definition ·resolved· to by the ·actual value· of the base [attribute] on the <restriction> or
			// <extension> element appearing as a child of <simpleContent>
			typeDef.baseTypeDefinition = g.resolveType(s, s.resolveQName(node.SimpleContent.Extension.Base))
			typeDef.derivationMethod = "extension"
		}

		typeDef.contentType.variety = "simple"

		if baseDef, ok := typeDef.baseTypeDefinition.(*complexTypeDefinition); ok {
			if baseDef.contentType.variety == "simple" {
				if baseDef.derivationMethod == "restriction" {
					// 1 If the {base type definition} is a complex type definition whose own {content type} has
					// {variety} simple and the <restriction> alternative is chosen, then let B be:
					//
					// 1.1 the simple type definition corresponding to the <simpleType> among the [children] of
					//     <restriction> if there is one;
					// 1.2 otherwise (<restriction> has no <simpleType> among its [children]), the simple type
					//     definition which is the {simple type definition} of the {content type} of the
					//     {base type definition}
					//
					// a simple type definition as follows:
				} else {
					// 3 If the {base type definition} is a complex type definition whose own {content type} has
					// {variety} simple and the <extension> alternative is chosen, then the {simple type definition} of
					// the {content type} of that complex type definition;
					typeDef.contentType.simpleTypeDefinition = baseDef.contentType.simpleTypeDefinition
				}
			} else if baseDef.contentType.variety == "mixed" {
				// 2 If the {base type definition} is a complex type definition whose own {content type} has {variety}
				// mixed and {particle} a Particle which is ·emptiable·, as defined in Particle Emptiable (§3.9.6.3) and
				// the <restriction> alternative is chosen, then (let SB be the simple type definition corresponding to
				// the <simpleType> among the [children] of <restriction> if any, otherwise ·xs:anySimpleType·) a simple
				// type definition which restricts SB with a set of constrainingFacet components corresponding to the appropriate
				// element information items among the <restriction>'s [children] (i.e. those which specify facets,
				// if any), as defined in Simple Type Restriction (Facets) (§3.16.6.4);
				//
				// Note: If there is no <simpleType> among the [children] of <restriction> (and if therefore SB is
				// ·xs:anySimpleType·), the result will be a simple type definition component which fails to obey the
				// constraints on simple type definitions, including for example clause 1.1 of Derivation Valid
				// (Restriction, Simple) (§3.16.6.2).
			}
		}
		if baseDef, ok := typeDef.baseTypeDefinition.(*simpleTypeDefinition); ok {
			// 4 If the {base type definition} is a simple type definition and the <extension> alternative is chosen,
			// then that simple type definition;
			if typeDef.derivationMethod == "extension" {
				typeDef.contentType.simpleTypeDefinition = baseDef
			}
		}
		if typeDef.contentType.simpleTypeDefinition == nil {
			typeDef.contentType.simpleTypeDefinition = anySimpleType
		}
	} else {
		if node.ComplexContent != nil {
			// 3.4.2.3.1 Mapping Rules for Complex Types with Explicit Complex Content

			// If the <restriction> alternative is chosen, then restriction, otherwise (the <extension> alternative is
			// chosen) extension.
			if node.ComplexContent.Restriction != nil {
				// The type definition ·resolved· to by the ·actual value· of the base [attribute] on the <restriction> or
				// <extension> element appearing as a child of <simpleContent>
				typeDef.baseTypeDefinition = g.resolveType(s, s.resolveQName(node.ComplexContent.Restriction.Base))
				typeDef.derivationMethod = "restriction"
			} else {
				// The type definition ·resolved· to by the ·actual value· of the base [attribute] on the <restriction> or
				// <extension> element appearing as a child of <simpleContent>
				typeDef.baseTypeDefinition = g.resolveType(s, s.resolveQName(node.ComplexContent.Extension.Base))
				typeDef.derivationMethod = "extension"
			}
		} else {
			// 3.4.2.3.2 Mapping Rules for Complex Types with Implicit Complex Content
			typeDef.baseTypeDefinition = nil
			typeDef.derivationMethod = "restriction"
		}

		// 3.4.2.3.3 Mapping Rules for Content Type Property of Complex Content
		//effectiveMixed := false
		//if node.ComplexContent.Mixed != nil {
		//	effectiveMixed = *node.ComplexContent.Mixed
		//} else if node.Mixed != nil {
		//	effectiveMixed = *node.Mixed
		//}

		//var explicitContent interface{}

	}

	return &typeDef
}

// resolveType resolves a qname into Type Definition
func (g *Generator) resolveType(s *schema, name xml.Name) TypeDefinition {
	// Check if type is a built-in type
	if typeDef, ok := xmlTypes[name]; ok {
		return typeDef
	}

	// Check if type is already parsed
	if typeDef, ok := s.typeDefinitions[name]; ok {
		return typeDef
	}

	// Find and parse type definition
	if s.targetNamespace == s.targetNamespace {
		for _, top := range s.xsdSchema.SchemaTop {
			switch t := top.(type) {
			case xsd.SimpleType:
				//
			case xsd.ComplexType:
				if t.Name == name.Local {
					return g.newComplexType(s, nil, &t)
				}
			}
		}
	}
	fmt.Println("Not found", name)
	return nil
}

func getBlocks(node *xsd.ComplexType, s *schema, typeDef complexTypeDefinition) []string {
	blocks := make([]string, 0)
	var effectiveBlockValue string
	if node.Block != "" {
		effectiveBlockValue = node.Block
	} else if s.blockDefault != "" {
		effectiveBlockValue = s.blockDefault
	} else {
		effectiveBlockValue = ""
	}
	if effectiveBlockValue == "#all" {
		blocks = []string{"extension", "restriction"}
	} else if effectiveBlockValue != "" {
		x := strings.Split(effectiveBlockValue, " ")
		for _, bl := range x {
			if bl == "extension" || bl == "restriction" {
				blocks = append(blocks, bl)
			}
		}
	}

	return blocks
}

func getFinals(node *xsd.ComplexType, s *schema, typeDef complexTypeDefinition) []string {
	blocks := make([]string, 0)
	var effectiveFinalValue string
	if node.Final != "" {
		effectiveFinalValue = node.Final
	} else if s.finalDefault != "" {
		effectiveFinalValue = s.finalDefault
	} else {
		effectiveFinalValue = ""
	}
	if effectiveFinalValue == "#all" {
		blocks = []string{"extension", "restriction"}
	} else if effectiveFinalValue != "" {
		x := strings.Split(effectiveFinalValue, " ")
		for _, bl := range x {
			if bl == "extension" || bl == "restriction" {
				blocks = append(blocks, bl)
			}
		}
	}

	return blocks
}

// 3.3.2.1 Common Mapping Rules for Element Declarations
func (g *Generator) newElement(s *schema, node xsd.Element) elementDeclaration {
	elm := elementDeclaration{}
	// The ·actual value· of the name [attribute].
	elm.name.Local = node.Name
	// The first of the following that applies:
	// 1 The type definition corresponding to the <simpleType> or <complexType> element information item in the
	//   [children], if either is present.
	// 2 The type definition ·resolved· to by the ·actual value· of the type [attribute], if it is present.
	// 3 The declared {type definition} of the Element Declaration ·resolved· to by the first QName in the
	//   ·actual value· of the substitutionGroup [attribute], if present.
	// 4 ·xs:anyType·.
	if node.ComplexType != nil {
		elm.typeDefinition = g.newComplexType(s, node, node.ComplexType)
	} else if node.SimpleType != nil {

	} else if node.Type != "" {
		elm.typeDefinition = g.resolveType(s, s.resolveQName(node.Type))
	} else {
		elm.typeDefinition = nil
	}
	// A Type Table corresponding to the <alternative> element information items among the [children], if any, as
	// follows, otherwise ·absent·.
	// elm.typeTable
	elm.nillable = node.Nillable
	// If there is a default or a fixed [attribute], then a Value Constraint as follows, otherwise ·absent·.
	// [Definition:]  Use the name effective simple type definition for the declared {type definition}, if it is
	// a simple type definition, or, if {type definition}.{content type}.{variety} = simple, for {type definition}.
	// {content type}.{simple type definition}, or else for the built-in string simple type definition).
	if node.Default != "" {
		actualValue := node.Default
		elm.valueConstraint = &valueConstraint{
			variety: "default",
			// the ·actual value· (with respect to the ·effective simple type definition·) of the [attribute]
			value: actualValue,
			// the ·normalized value· (with respect to the ·effective simple type definition·) of the [attribute]
			lexicalForm: normalizeValue(actualValue),
		}
	}
	if node.Fixed != "" {
		actualValue := node.Fixed
		elm.valueConstraint = &valueConstraint{
			variety: "default",
			// the ·actual value· (with respect to the ·effective simple type definition·) of the [attribute]
			value: actualValue,
			// the ·normalized value· (with respect to the ·effective simple type definition·) of the [attribute]
			lexicalForm: normalizeValue(actualValue),
		}
	}
	// A set consisting of the identity-constraint-definitions corresponding to all the <key>, <unique> and
	// <keyref> element information items in the [children], if any, otherwise the empty set.
	// elm.identityConstraintDefinitions
	// A set of the element declarations ·resolved· to by the items in the ·actual value· of the substitutionGroup
	// [attribute], if present, otherwise the empty set.
	// elm.substitutionGroupAffiliations
	// A set depending on the ·actual value· of the block [attribute], if present, otherwise on the ·actual value·
	// of the blockDefault [attribute] of the ancestor <schema> element information item, if present, otherwise
	// on the empty string. Call this the EBV (for effective block value). Then the value of this property is the
	// appropriate case among the following:
	//
	// 1 If the EBV is the empty string, then the empty set;
	// 2 If the EBV is #all, then {extension, restriction, substitution};
	// 3 otherwise a set with members drawn from the set above, each being present or absent depending on whether
	//   the ·actual value· (which is a list) contains an equivalently named item.
	//
	// Note: Although the blockDefault [attribute] of <schema> may include values other than extension, restriction
	// or substitution, those values are ignored in the determination of {disallowed substitutions} for element
	// declarations (they are used elsewhere).
	// elm.disallowedSubstitutions
	// As for {disallowed substitutions} above, but using the final and finalDefault [attributes] in place of the
	// block and blockDefault [attributes] and with the relevant set being {extension, restriction}.
	//elm.substitutionGroupExclusions
	// The ·actual value· of the abstract [attribute], if present, otherwise false.
	elm.abstract = node.Abstract
	// The ·annotation mapping· of the <element> element and any of its <unique>, <key> and <keyref> [children]
	// with a ref [attribute], as defined in XML Representation of Annotation Schema Components (§3.15.2).
	// elm.annotations
	return elm
}

func normalizeValue(s string) string {
	// replace
	r := regexp.MustCompile("[\t\r\n]").ReplaceAllString(s, " ")
	// collapse
	r = strings.Trim(regexp.MustCompile(" +").ReplaceAllString(r, " "), " ")
	return r
}
