package goxsd

import (
	"encoding/xml"
)

const xmlNs = "http://www.w3.org/2001/XMLSchema"

var xmlTypes = map[xml.Name]TypeDefinition{
	anyType.name:         anyType,
	anySimpleType.name:   anySimpleType,
	anyAtomicType.name:   anyAtomicType,
	stringPrimitive.name: stringPrimitive,
}

var anyType = &complexTypeDefinition{
	name:             xml.Name{Space: xmlNs, Local: "anyType"},
	derivationMethod: "restriction",
	contentType: complexTypeContentType{
		variety: "mixed",
		particle: &particle{
			minOccurs: 1,
			maxOccurs: 1,
			term: modelGroup{
				compositor: "sequence",
				particles: []particle{
					{
						minOccurs: 0,
						maxOccurs: unbound,
						term: wildcard{
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
	attributeUses: []attributeUse{},
	attributeWildcard: wildcard{
		namespaceConstraint: wildcardNamespaceConstraint{
			variety:         "any",
			namespaces:      []string{},
			disallowedNames: []string{},
		},
		processContents: "lax",
	},
	final: []string{},
	prohibitedSubstitutions: []string{},
	assertions:              []assertion{},
	annotations:             []annotation{},
	abstract:                false,
}

var anySimpleType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "anySimpleType"},
	final:              []string{},
	facets:             []ConstrainingFacet{},
	baseTypeDefinition: anyType,
	fundamentalFacets:  []FundamentalFacet{},
	annotations:        []annotation{},
}

var anyAtomicType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "anyAtomicType"},
	final:              []string{},
	baseTypeDefinition: anySimpleType,
	facets:             []ConstrainingFacet{},
	variety:            "atomic",
	fundamentalFacets:  []FundamentalFacet{},
	annotations:        []annotation{},
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

func init() {
	stringPrimitive.goType = "string"
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
		annotations:        []annotation{},
	}
	t.primitiveTypeDefinition = t
	return t
}
