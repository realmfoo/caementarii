package goxsd

type attrId struct {
	Id string `xml:"id,attr"`
}

type attrFixed struct {
	Fixed bool `xml:"fixed,attr"`
}

type attrName struct {
	Name NCName `xml:"name,attr"`
}

type attrBase struct {
	Base QName `xml:"base,attr"`
}

type hasAnnotation struct {
	Annotation *XmlAnnotation `xml:"annotation,omitempty"`
}

type XmlAttribute struct {
	Default string `xml:"default,attr"`
	Fixed   string `xml:"fixed,attr"`
	Form    string `xml:"form,attr"`
	attrId
	attrName
	Ref             QName  `xml:"ref,attr"`
	TargetNamespace anyURI `xml:"targetNamespace,attr"`
	Type            QName  `xml:"type,attr"`
	Use             string `xml:"use,attr"`
	Inheritable     bool   `xml:"inheritable,attr"`

	Annotation *XmlAnnotation `xml:"annotation,omitempty"`
	SimpleType *XmlSimpleType `xml:"simpleType,omitempty"`
}

type XmlAnnotation struct {
	attrId

	AppInfo       *XmlAppInfo       `xml:"appinfo"`
	Documentation *XmlDocumentation `xml:"documentation"`
}

type XmlAppInfo struct {
	Source  anyURI `xml:"source,attr"`
	Content string `xml:",chardata"`
}

type XmlDocumentation struct {
	Source  anyURI `xml:"source,attr"`
	Content string `xml:",chardata"`
}

type XmlMinExclusive struct {
	attrFixed
	attrId
	Value string `xml:"value,attr"`

	hasAnnotation
}

type XmlMinInclusive struct {
	attrFixed
	attrId
	Value string `xml:"value,attr"`

	hasAnnotation
}

type XmlMaxExclusive struct {
	attrFixed
	attrId
	Value string `xml:"value,attr"`

	hasAnnotation
}

type XmlMaxInclusive struct {
	attrFixed
	attrId
	Value string `xml:"value,attr"`

	hasAnnotation
}

type XmlTotalDigits struct {
	attrFixed
	attrId
	Value int `xml:"value,attr"`

	hasAnnotation
}

type XmlFractionDigits struct {
	attrFixed
	attrId
	Value int `xml:"value,attr"`

	hasAnnotation
}

type XmlLength struct {
	attrFixed
	attrId
	Value int `xml:"value,attr"`

	hasAnnotation
}

type XmlMinLength struct {
	attrFixed
	attrId
	Value int `xml:"value,attr"`

	hasAnnotation
}

type XmlMaxLength struct {
	attrFixed
	attrId
	Value int `xml:"value,attr"`

	hasAnnotation
}

type XmlEnumeration struct {
	attrId
	Value string `xml:"value,attr"`

	hasAnnotation
}

type XmlWhiteSpace struct {
	attrFixed
	attrId
	// (collapse | preserve | replace)
	Value string `xml:"value,attr"`

	hasAnnotation
}

type XmlPattern struct {
	attrId
	Value string `xml:"value,attr"`

	hasAnnotation
}

type XmlAssertion struct {
	attrId
	Test                  string `xml:"test,attr"`
	XPathDefaultNamespace string `xml:"xpathDefaultNamespace,attr"`

	hasAnnotation
}

type XmlExplicitTimezone struct {
	attrFixed
	attrId
	Value NCName `xml:"value,attr"`

	hasAnnotation
}

type XmlSimpleType struct {
	Final string `xml:"final,attr"`
	attrId
	attrName

	hasAnnotation
	Restriction *struct {
		attrBase
		attrId

		hasAnnotation
		SimpleType       []XmlSimpleType       `xml:"simpleType,omitempty"`
		MinExclusive     []XmlMinExclusive     `xml:"minExclusive,omitempty"`
		MinInclusive     []XmlMinInclusive     `xml:"minInclusive,omitempty"`
		MaxExclusive     []XmlMaxExclusive     `xml:"maxExclusive,omitempty"`
		MaxInclusive     []XmlMaxInclusive     `xml:"maxInclusive,omitempty"`
		TotalDigits      []XmlTotalDigits      `xml:"totalDigits,omitempty"`
		FractionDigits   []XmlFractionDigits   `xml:"fractionDigits,omitempty"`
		Length           []XmlLength           `xml:"length,omitempty"`
		MinLength        []XmlMinLength        `xml:"minLength,omitempty"`
		MaxLength        []XmlMaxLength        `xml:"maxLength,omitempty"`
		Enumeration      []XmlEnumeration      `xml:"enumeration,omitempty"`
		WhiteSpace       []XmlWhiteSpace       `xml:"whiteSpace,omitempty"`
		Pattern          []XmlPattern          `xml:"pattern,omitempty"`
		Assertion        []XmlAssertion        `xml:"assertion,omitempty"`
		ExplicitTimezone []XmlExplicitTimezone `xml:"explicitTimezone,omitempty"`
	} `xml:"restriction,omitempty"`
	List *struct {
		attrId
		ItemType QName `xml:"itemType,attr"`

		hasAnnotation
		SimpleType *XmlSimpleType `xml:"simpleType,omitempty"`
	} `xml:"list,omitempty"`
	Union *struct {
		attrId
		MemberTypes ListOfQName `xml:"memberTypes,attr"`

		hasAnnotation
		SimpleTypes []*XmlSimpleType `xml:"simpleType,omitempty"`
	} `xml:"union,omitempty"`
}
