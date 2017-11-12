package goxsd

import (
	"encoding/xml"
	"github.com/realmfoo/caementarii/xsd"
	"math"
	"strings"
)

var unbounded = math.MaxInt32

type TypeDefinition interface {
	aTypeDef()
}

type typeDefinition struct {
}

func (*typeDefinition) aTypeDef() {}

// Attribute declarations provide for:
//
// * Local ·validation· of attribute information item values using a simple type definition;
// * Specifying default or fixed values for attribute information items.
type attributeDeclaration struct {
	// A name with optional target namespace.
	name xml.Name
	// A Simple Type Definition component. Required.
	typeDefinition *simpleTypeDefinition
	scope          struct {
		// One of {global, local}. Required.
		variety string
		// Either a Complex Type Definition or a Attribute Group Definition. Required if {variety} is local, otherwise must be ·absent·
		parent interface{}
	}
	valueConstraint valueConstraint
	inheritable     bool

	annotatedComponent
}

// Element declarations provide for:
//
// * Local ·validation· of element information item values using a type definition;
// * Specifying default or fixed values for element information items;
// * Establishing uniquenesses and reference constraint relationships among the values of related elements and attributes;
// * Controlling the substitutability of elements through the mechanism of ·element substitution groups·.
type elementDeclaration struct {
	// A name with optional target namespace.
	name xml.Name
	// A Type Definition component. Required.
	typeDefinition interface{}
	typeTable      *struct {
		// A sequence of Type Alternative components.
		alternatives []alternative
		// A Type Alternative component. Required.
		defaultTypeDefinition interface{}
	}
	scope struct {
		// One of {global, local}. Required.
		variety string
		// Either a Complex Type Definition or a Model Group Definition. Required if {variety} is local, otherwise must be ·absent·
		parent interface{}
	}
	valueConstraint *valueConstraint
	// An xs:boolean value. Required.
	nillable bool
	// A set of Identity-Constraint Definition components.
	identityConstraintDefinitions []identityConstraint
	// A set of Element Declaration components.
	substitutionGroupAffiliations []elementDeclaration
	// A subset of {extension, restriction}.
	substitutionGroupExclusions []string
	// A subset of {substitution, extension, restriction}.
	disallowedSubstitutions []string
	// An xs:boolean value. Required.
	abstract bool

	annotatedComponent
}

type alternative struct {
	// An XPath Expression property record. Optional.
	test *xpathExpression
	// A Type Definition component. Required.
	typeDefinition TypeDefinition

	annotatedComponent
}

// XPath Expression
type xpathExpression struct {
	// A set of Namespace Binding property records.
	namespaceBindings []xml.Name
	// An xs:anyURI value. Optional.
	defaultNamespace string
	// An xs:anyURI value. Optional.
	baseURI string
	// An [XPath 2.0] expression. Required.
	expression string
}

// Identity-constraint definition components provide for uniqueness and reference constraints with respect to the contents of multiple elements and attributes.
type identityConstraint struct {
	// A name with optional target namespace.
	name xml.Name
	// One of {key, keyref, unique}. Required.
	identityConstraintCategory string
	// An XPath Expression property record. Required.
	selector xpathExpression
	// A sequence of XPath Expression property records.
	fields []xpathExpression
	// An Identity-Constraint Definition component. Required if {identity-constraint category} is keyref, otherwise ({identity-constraint category} is key or unique) must be ·absent·.
	// If a value is present, its {identity-constraint category} must be key or unique.
	referencedKey *identityConstraint

	annotatedComponent
}

type annotation struct {
	// A sequence of Element information items.
	applicationInformation []string
	// A sequence of Element information items.
	userInformation []string
	// A set of Attribute information items.
	attributes []xml.Attr
}

// Complex Type Definition, a kind of Type Definition
type complexTypeDefinition struct {
	// A name with optional target namespace.
	name xml.Name
	// A Type Definition component. Required.
	baseTypeDefinition TypeDefinition
	// A subset of {extension, restriction}.
	final []string
	// Required if {name} is ·absent·, otherwise must be ·absent·.
	// Either an Element Declaration or a Complex Type Definition.
	context interface{}
	// One of {extension, restriction}. Required.
	derivationMethod string
	// An xs:boolean value. Required.
	abstract bool
	// A set of Attribute Use components.
	attributeUses []*attributeUse
	// A Wildcard component. Optional.
	attributeWildcard wildcard
	// A Content Type property record. Required.
	contentType complexTypeContentType
	// A subset of {extension, restriction}.
	prohibitedSubstitutions []string
	// A sequence of Assertion components.
	assertions []assertion

	typeDefinition
	annotatedComponent
}

type complexTypeContentType struct {
	// One of {empty, simple, element-only, mixed}. Required.
	variety string
	// A Particle component. Required if {variety} is element-only or mixed, otherwise must be ·absent·.
	particle *particle
	// An Open Content property record. Optional if {variety} is element-only or mixed, otherwise must be ·absent·.
	openContent *openContent
	// A Simple Type Definition component. Required if {variety} is simple, otherwise must be ·absent·.
	simpleTypeDefinition *simpleTypeDefinition
}

type openContent struct {
	// One of {interleave, suffix}. Required.
	mode string
	// A Wildcard component. Required.
	wildcard wildcard
}

// Simple Type Definition, a kind of Type Definition
type simpleTypeDefinition struct {
	// A name with optional target namespace.
	name xml.Name
	// A subset of {extension, restriction, list, union}.
	final []string
	// Required if {name} is ·absent·, otherwise must be ·absent·.
	// Either an Attribute Declaration, an Element Declaration, a Complex Type Definition, or a Simple Type Definition.
	context interface{}
	// A Type Definition component. Required.
	// With one exception, the {base type definition} of any Simple Type Definition is a Simple Type Definition. The exception is ·xs:anySimpleType·, which has ·xs:anyType·, a Complex Type Definition, as its {base type definition}.
	baseTypeDefinition TypeDefinition
	// A set of Constraining Facet components.
	facets []ConstrainingFacet
	// A set of Fundamental Facet components.
	fundamentalFacets []FundamentalFacet
	// One of {atomic, list, union}. Required for all Simple Type Definitions except ·xs:anySimpleType·, in which it is ·absent·.
	variety string
	// A Simple Type Definition component. With one exception, required if {variety} is atomic, otherwise must be ·absent·. The exception is ·xs:anyAtomicType·, whose {primitive type definition} is ·absent·.
	// If non-·absent·, must be a primitive definition.
	primitiveTypeDefinition *simpleTypeDefinition
	// A Simple Type Definition component. Required if {variety} is list, otherwise must be ·absent·.
	// The value of this property must be a primitive or ordinary simple type definition with {variety} = atomic, or an ordinary simple type definition with {variety} = union whose basic members are all atomic; the value must not itself be a list type (have {variety} = list) or have any basic members which are list types.
	itemTypeDefinition *simpleTypeDefinition
	// A sequence of primitive or ordinary Simple Type Definition components.
	// Must be present (but may be empty) if {variety} is union, otherwise must be ·absent·.
	// The sequence may contain any primitive or ordinary simple type definition, but must not contain any special type definitions.
	numberTypeDefinitions []TypeDefinition

	// A Go type for representing a content
	goType string

	typeDefinition
	annotatedComponent
}

type assertion struct {
	// An XPath Expression property record. Required.
	test xpathExpression

	annotatedComponent
}

// An attribute use is a utility component which controls the occurrence and defaulting behavior of attribute
// declarations. It plays the same role for attribute declarations in complex types that particles play for element
// declarations.
type attributeUse struct {
	// An xs:boolean value. Required.
	required bool
	// An Attribute Declaration component. Required.
	attributeDeclaration *attributeDeclaration
	// A Value Constraint property record. Optional.
	valueConstraint *valueConstraint
	// An xs:boolean value. Required.
	inheritable bool

	annotatedComponent
}

// A schema can name a group of attribute declarations so that they can be incorporated as a group into complex type definitions.
// Attribute group definitions do not participate in ·validation· as such, but the {attribute uses} and {attribute wildcard} of one or more complex type definitions may be constructed in whole or part by reference to an attribute group. Thus, attribute group definitions provide a replacement for some uses of XML's parameter entity facility. Attribute group definitions are provided primarily for reference from the XML representation of schema components (see <complexType> and <attributeGroup>).
type attributeGroupDefinition struct {
	// A name with optional target namespace.
	name xml.Name
	// A set of Attribute Use components.
	attributeUses []attributeUse
	// A Wildcard component. Optional.
	attributeWildcard wildcard

	annotatedComponent
}

// In order to exploit the full potential for extensibility offered by XML plus namespaces, more provision is needed than DTDs allow for targeted flexibility in content models and attribute declarations. A wildcard provides for ·validation· of attribute and element information items dependent on their namespace names and optionally on their local names.
type wildcard struct {
	// A Namespace Constraint property record. Required.
	namespaceConstraint wildcardNamespaceConstraint
	// One of {skip, strict, lax}. Required.
	processContents string

	annotatedComponent
}

type wildcardNamespaceConstraint struct {
	// One of {any, enumeration, not}. Required.
	variety string
	// A set each of whose members is either an xs:anyURI value or the distinguished value ·absent·. Required.
	namespaces []string
	// A set each of whose members is either an xs:QName value or the keyword defined or the keyword sibling. Required.
	disallowedNames []string
}

// A model group definition associates a name and optional annotations with a Model Group. By reference to the name, the entire model group can be incorporated by reference into a {term}.
// Model group definitions are provided primarily for reference from the XML Representation of Complex Type Definition Schema Components (§3.4.2) (see <complexType> and <group>). Thus, model group definitions provide a replacement for some uses of XML's parameter entity facility.
type modelGroupDefinition struct {
	// A name with optional target namespace.
	name xml.Name
	// A Model Group component. Required.
	modelGroup modelGroup

	annotatedComponent
}

type modelGroup struct {
	// One of {all, choice, sequence}. Required.
	compositor string
	// A sequence of Particle components.
	particles []*particle

	annotatedComponent
}

type particle struct {
	// An xs:nonNegativeInteger value. Required.
	minOccurs int
	// Either a positive integer or unbounded. Required.
	maxOccurs int
	// A Term component. Required.
	term AnnotatedComponent

	annotatedComponent
}

// Notation declarations reconstruct XML NOTATION declarations.
type notationDeclaration struct {
	// A name with optional target namespace.
	name xml.Name
	// An xs:anyURI value. Required if {public identifier} is ·absent·, otherwise ({public identifier} is present) optional.
	systemIdentifier string
	// A publicID value. Required if {system identifier} is ·absent·, otherwise ({system identifier} is present) optional.
	// As defined in [XML 1.0] or [XML 1.1].
	publicIdentifier string

	annotatedComponent
}

// At the abstract level, the schema itself is just a container for its components.
type schema struct {
	// A sequence of Annotation components.
	annotations []annotation
	// A set of Type Definition components.
	typeDefinitions map[xml.Name]TypeDefinition
	// A set of Attribute Declaration components.
	attributeDeclarations []attributeDeclaration
	// A set of Element Declaration components.
	elementDeclarations map[xml.Name]*elementDeclaration
	// A set of Attribute Group Definition components.
	attributeGroupDefinitions []attributeGroupDefinition
	// A set of Model Group Definition components.
	modelGroupDefinitions []modelGroupDefinition
	// A set of Notation Declaration components.
	notationDeclarations []notationDeclaration
	// A set of Identity-Constraint Definition components.
	identityConstraintDefinitions []identityConstraint

	xsdSchema *xsd.Schema
	// A map of known namespaces
	prefixMap            map[string]string
	targetNamespace      string
	blockDefault         string
	finalDefault         string
	attributeFormDefault string
}

func newSchema(s *xsd.Schema) *schema {
	prefixMap := make(map[string]string, 0)
	for _, attr := range s.XMLAttrs {
		if attr.Name.Space == "xmlns" {
			prefixMap[attr.Name.Local] = attr.Value
		}
	}
	return &schema{
		xsdSchema:           s,
		targetNamespace:     s.TargetNamespace,
		prefixMap:           prefixMap,
		typeDefinitions:     make(map[xml.Name]TypeDefinition, 0),
		elementDeclarations: make(map[xml.Name]*elementDeclaration, 0),
	}
}

// resolveQName resolves a QName value into xml.Name struct
func (s *schema) resolveQName(qname string) (name xml.Name) {
	p := strings.SplitN(qname, ":", 2)
	if len(p) == 1 {
		name.Local = p[0]
	} else {
		name.Space = s.prefixMap[p[0]]
		name.Local = p[1]
	}
	return
}

//-----------------------------------------------------------------------------
// Constraining Facets

type ConstrainingFacet interface {
	aConstrainingFacet()
}

type (
	valueConstraint struct {
		// One of {default, fixed}. Required.
		variety string
		// An actual value. Required
		value interface{}
		// A character string. Required.
		lexicalForm string

		constrainingFacet
	}

	whiteSpaceFacet struct {
		// A sequence of Annotation components.
		annotations []annotation
		// One of {preserve, replace, collapse}. Required.
		value string
		// An xs:boolean value. Required.
		fixed bool

		constrainingFacet
	}
)

type constrainingFacet struct{}

func (*constrainingFacet) aConstrainingFacet() {}

//-----------------------------------------------------------------------------
// Fundamental Facets

type FundamentalFacet interface {
	aFundamentalFacet()
}

type (
	// One of {false, partial, total}. Required.
	orderedFacet struct {
		string
		fundamentalFacet
	}

	boundedFacet struct {
		bool
		fundamentalFacet
	}

	// One of {finite, countably infinite}.
	cardinalityFacet struct {
		string
		fundamentalFacet
	}

	numericFacet struct {
		bool
		fundamentalFacet
	}
)

type fundamentalFacet struct{}

func (*fundamentalFacet) aFundamentalFacet() {}

//-----------------------------------------------------------------------------
// Annotated Component

type AnnotatedComponent interface {
	anAnnotatedComponent()
}

type annotatedComponent struct {
	// A sequence of Annotation components.
	annotations []annotation
}

func (*annotatedComponent) anAnnotatedComponent() {}
