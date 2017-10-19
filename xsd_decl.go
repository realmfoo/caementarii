package goxsd

import "github.com/realmfoo/caementarii/xsd"

type attributeDecl struct {
	// A sequence of annotation components
	annotations []Annotation

	// An xs:NCName value. Required.
	name xsd.NCName

	// An xs:anyURI value. Optional.
	targetNamespace string

	// A Simple Type Definition component. Required.
	//	typeDefinition SimpleTypeDef

	// A Scope property record. Required.
	scope struct {
		// One of {global, local}. Required.
		variety string

		// Either a Complex Type Definition or a Attribute Group Definition. Required if {variety} is local, otherwise must be ·absent·
		parent *interface{}
	}

	// A Value Constraint property record. Optional.
	valueConstraint struct {
		// One of {default, fixed}. Required.
		variety string

		// An ·actual value·. Required.
		value string

		// A character string. Required.
		lexicalForm string
	}

	// An xs:boolean value. Required.
	inheritable bool
}

// Annotation is a kind of component
type Annotation struct {
	// Application information
	// User information
	// Attributes
}
