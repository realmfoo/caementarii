package goxsd

import (
	"bufio"
	"fmt"
	"github.com/realmfoo/caementarii/xsd"
	"os"
	"regexp"
	"strings"
)

type Generator struct {
	pkgName string
}

func (g *Generator) generate(s *xsd.Schema) {

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	schema := &schema{}
	schema.targetNamespace = s.TargetNamespace

	for _, top := range s.SchemaTop {
		if node, ok := top.(xsd.Element); ok {
			// 3.3.2.1 Common Mapping Rules for Element Declarations
			elm := g.newElement(schema, node)

			// 3.3.2.2 Mapping Rules for Top-Level Element Declarations
			elm.name.Space = s.TargetNamespace

			elm.scope.variety = "global"

			fmt.Printf("%+v\n", elm)
			schema.elementDeclarations = append(schema.elementDeclarations, elm)
		}
	}

	w.WriteString("package " + g.pkgName + "\n\n")
	w.WriteString("import (\n")
	w.WriteString(")\n\n")
}

func (g *Generator) newComplexType(s *schema, parent interface{}, node *xsd.ComplexType) complexTypeDefinition {
	typeDef := complexTypeDefinition{}

	// The ·actual value· of the name [attribute].
	typeDef.name.Local = node.Name
	// The ·actual value· of the targetNamespace [attribute] of the <schema> ancestor element information item if present, otherwise ·absent·.
	typeDef.name.Space = s.targetNamespace
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
	//typeDef.assertions =

	// The ·annotation mapping· of the set of elements containing the <complexType>, the <openContent> [child], if
	// present, the <attributeGroup> [children], if present, the <simpleContent> and <complexContent> [children], if
	// present, and their <restriction> and <extension> [children], if present, and their <openContent> and
	// <attributeGroup> [children], if present, as defined in
	// XML Representation of Annotation Schema Components (§3.15.2).
	//typeDef.annotations =

	if node.SimpleContent != nil {

	} else {
		if node.ComplexContent != nil {
			// 3.4.2.3.1 Mapping Rules for Complex Types with Explicit Complex Content

			// The type definition ·resolved· to by the ·actual value· of the base [attribute]
			// typeDef.baseTypeDefinition =

			// If the <restriction> alternative is chosen, then restriction, otherwise (the <extension> alternative is
			// chosen) extension.
			if node.ComplexContent.Restriction != nil {
				typeDef.derivationMethod = "restriction"
			} else {
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

	return typeDef
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
