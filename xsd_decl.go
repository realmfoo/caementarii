package goxsd

// NCName represents XML "non-colonized" Names.
// white-space: collapse
type NCName string

// QName represents XML qualified names.
type QName string

// A list of QName
type ListOfQName []QName

// anyURI represents an Internationalized Resource Identifier Reference (IRI).
type anyURI string

type attributeDecl struct {
	// A sequence of annotation components
	annotations []Annotation

	// An xs:NCName value. Required.
	name NCName

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
