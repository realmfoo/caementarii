package goxsd

import (
	"encoding/xml"
)

const xmlNs = "http://www.w3.org/2001/XMLSchema"

var xmlTypes = map[xml.Name]TypeDefinition{
	anyType.name:          anyType,
	anySimpleType.name:    anySimpleType,
	anyAtomicType.name:    anyAtomicType,
	stringPrimitive.name:  stringPrimitive,
	decimalPrimitive.name: decimalPrimitive,
	booleanPrimitive.name: booleanPrimitive,
	anyURIPrimitive.name:  anyURIPrimitive,
	qNamePrimitive.name:   qNamePrimitive,

	normalizedStringDataType.name: normalizedStringDataType,
	tokenDataType.name:            tokenDataType,
	nmTokenDataType.name:          nmTokenDataType,
	nameDataType.name:             nameDataType,
	ncNameDataType.name:           ncNameDataType,
	idDataType.name:               idDataType,
	integerType.name:              integerType,
	nonNegativeIntegerType.name:   nonNegativeIntegerType,
	positiveIntegerType.name:      positiveIntegerType,
}

var anyType = &complexTypeDefinition{
	name:             xml.Name{Space: xmlNs, Local: "anyType"},
	derivationMethod: "restriction",
	contentType: complexTypeContentType{
		variety: "mixed",
		particle: &particle{
			minOccurs: 1,
			maxOccurs: 1,
			term: &modelGroup{
				compositor: "sequence",
				particles: []*particle{
					{
						minOccurs: 0,
						maxOccurs: unbounded,
						term: &wildcard{
							namespaceConstraint: wildcardNamespaceConstraint{
								variety:         "any",
								namespaces:      []string{},
								disallowedNames: []string{},
							},
							processContents: "lax",
						},
					},
				},
			},
		},
	},
	attributeUses: []*attributeUse{},
	attributeWildcard: wildcard{
		namespaceConstraint: wildcardNamespaceConstraint{
			variety:         "any",
			namespaces:      []string{},
			disallowedNames: []string{},
		},
		processContents: "lax",
	},
	final:                   []string{},
	prohibitedSubstitutions: []string{},
	assertions:              []assertion{},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
	abstract: false,
}

var anySimpleType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "anySimpleType"},
	final:              []string{},
	facets:             []ConstrainingFacet{},
	baseTypeDefinition: anyType,
	fundamentalFacets:  []FundamentalFacet{},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
}

var anyAtomicType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "anyAtomicType"},
	final:              []string{},
	baseTypeDefinition: anySimpleType,
	facets:             []ConstrainingFacet{},
	variety:            "atomic",
	fundamentalFacets:  []FundamentalFacet{},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
}

var stringPrimitive = newPrimitive(
	"string",
	[]ConstrainingFacet{
		&whiteSpaceFacet{value: "preserve", fixed: false},
	},
	[]FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
)

var decimalPrimitive = newPrimitive(
	"decimal",
	[]ConstrainingFacet{
		&whiteSpaceFacet{value: "collapse", fixed: false},
	},
	[]FundamentalFacet{
		&orderedFacet{string: "total"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: true},
	},
)

var integerType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "integer"},
	baseTypeDefinition: decimalPrimitive,
	final:              []string{},
	variety:            "atomic",
	facets: []ConstrainingFacet{
		&whiteSpaceFacet{value: "collapse", fixed: true},
		&fractionDigitsFacet{value: 0, fixed: true},
	},
	fundamentalFacets: []FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
	goType: "int",
}

var nonNegativeIntegerType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "nonNegativeInteger"},
	baseTypeDefinition: integerType,
	final:              []string{},
	variety:            "atomic",
	facets: []ConstrainingFacet{
		&whiteSpaceFacet{value: "collapse", fixed: true},
		&fractionDigitsFacet{value: 0, fixed: true},
		&minInclusiveFacet{value: "0"},
	},
	fundamentalFacets: []FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
	goType: "uint",
}

var positiveIntegerType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "positiveInteger"},
	baseTypeDefinition: integerType,
	final:              []string{},
	variety:            "atomic",
	facets: []ConstrainingFacet{
		&whiteSpaceFacet{value: "collapse", fixed: true},
		&fractionDigitsFacet{value: 0, fixed: true},
		&minInclusiveFacet{value: "1"},
	},
	fundamentalFacets: []FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
	goType: "uint",
}

var booleanPrimitive = newPrimitive(
	"boolean",
	[]ConstrainingFacet{
		&patternFacet{},
		&whiteSpaceFacet{value: "preserve", fixed: false},
	},
	[]FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "finite"},
		&numericFacet{bool: false},
	},
)

var anyURIPrimitive = newPrimitive(
	"anyURI",
	[]ConstrainingFacet{
		&whiteSpaceFacet{value: "collapse", fixed: true},
	},
	[]FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
)

var qNamePrimitive = newPrimitive(
	"QName",
	[]ConstrainingFacet{
		&whiteSpaceFacet{value: "collapse", fixed: true},
	},
	[]FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
)

var normalizedStringDataType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "normalizedString"},
	baseTypeDefinition: stringPrimitive,
	final:              []string{},
	variety:            "atomic",
	facets: []ConstrainingFacet{
		&whiteSpaceFacet{value: "replace"},
	},
	fundamentalFacets: []FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
	goType: "string",
}

var tokenDataType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "token"},
	baseTypeDefinition: normalizedStringDataType,
	final:              []string{},
	variety:            "atomic",
	facets: []ConstrainingFacet{
		&whiteSpaceFacet{value: "collapse"},
	},
	fundamentalFacets: []FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
	goType: "string",
}

var nmTokenDataType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "NMTOKEN"},
	baseTypeDefinition: tokenDataType,
	final:              []string{},
	variety:            "atomic",
	facets: []ConstrainingFacet{
		&patternFacet{value: "\\c+"},
		&whiteSpaceFacet{value: "collapse"},
	},
	fundamentalFacets: []FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
	goType: "string",
}

var nameDataType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "Name"},
	baseTypeDefinition: tokenDataType,
	final:              []string{},
	variety:            "atomic",
	facets: []ConstrainingFacet{
		&whiteSpaceFacet{value: "collapse"},
	},
	fundamentalFacets: []FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
	goType: "string",
}

var ncNameDataType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "NCName"},
	baseTypeDefinition: nameDataType,
	final:              []string{},
	variety:            "atomic",
	facets: []ConstrainingFacet{
		&whiteSpaceFacet{value: "collapse"},
	},
	fundamentalFacets: []FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
	goType: "string",
}

var idDataType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "ID"},
	baseTypeDefinition: ncNameDataType,
	final:              []string{},
	variety:            "atomic",
	facets: []ConstrainingFacet{
		&whiteSpaceFacet{value: "collapse"},
	},
	fundamentalFacets: []FundamentalFacet{
		&orderedFacet{string: "false"},
		&boundedFacet{bool: false},
		&cardinalityFacet{string: "countably infinite"},
		&numericFacet{bool: false},
	},
	annotatedComponent: annotatedComponent{
		annotations: []annotation{},
	},
	goType: "string",
}

func init() {
	stringPrimitive.goType = "string"
	decimalPrimitive.goType = "float64"
	anyURIPrimitive.goType = "string"
	qNamePrimitive.goType = "string"
	booleanPrimitive.goType = "bool"
}

// newPrimitive creates a new primitive type by a template.
func newPrimitive(name string, facets []ConstrainingFacet, fundamentalFacets []FundamentalFacet) *simpleTypeDefinition {
	t := &simpleTypeDefinition{
		name:               xml.Name{Space: xmlNs, Local: name},
		baseTypeDefinition: anyAtomicType,
		final:              []string{},
		variety:            "atomic",
		facets:             facets,
		fundamentalFacets:  fundamentalFacets,
		annotatedComponent: annotatedComponent{
			annotations: []annotation{},
		},
	}
	t.primitiveTypeDefinition = t
	return t
}
