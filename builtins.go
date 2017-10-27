package goxsd

import (
	"encoding/xml"
	"math"
)

const xmlNs = "http://www.w3.org/2001/XMLSchema"

var xmlTypes = map[xml.Name]interface{}{
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
	facets:             []interface{}{},
	baseTypeDefinition: anyType,
	fundamentalFacets:  []interface{}{},
	annotations:        []annotation{},
}

var anyAtomicType = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "anyAtomicType"},
	final:              []string{},
	baseTypeDefinition: anySimpleType,
	facets:             []interface{}{},
	variety:            "atomic",
	fundamentalFacets:  []interface{}{},
	annotations:        []annotation{},
}

var stringPrimitive = &simpleTypeDefinition{
	name:               xml.Name{Space: xmlNs, Local: "string"},
	baseTypeDefinition: anyAtomicType,
	final:              []string{},
	variety:            "atomic",
	facets: []interface{}{
		whiteSpaceFacet{value: "preserve", fixed: false},
	},
	fundamentalFacets: []interface{}{
		orderedFacet("false"),
		boundedFacet(false),
		cardinalityFacet("countably infinite"),
		numericFacet(false),
	},
	annotations: []annotation{},
}

var unbound = math.MaxInt32
