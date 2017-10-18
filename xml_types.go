package goxsd

type XMLCompositions struct {
	Include    []XMLInclude    `xml:"include"`
	Import     []XMLImport     `xml:"import"`
	Redefine   []XMLRedefine   `xml:"redefine"`
	Override   []XMLOverride   `xml:"override"`
	Annotation []XMLAnnotation `xml:"annotation"`
}

// This group is for the elements which can self-redefine
type XMLRedefinables struct {
	SimpleType     []XMLSimpleType     `xml:"simpleType"`
	ComplexType    []XMLComplexType    `xml:"complexType"`
	Group          []XMLGroup          `xml:"group"`
	AttributeGroup []XMLAttributeGroup `xml:"attributeGroup"`
}

// This group is for the elements which occur freely at the top level of schemas.
// All of their types are based on the "annotated" type by extension.
type XMLSchemaTops struct {
	XMLRedefinables
	Element   []XMLElement   `xml:"element"`
	Attribute []XMLAttribute `xml:"attribute"`
	Notation  []XMLNotation  `xml:"notation"`
}

// {qualified, unqualified}
type formChoice string

// {extension, restriction}
type reducedDerivationControl string

// {#all} or (possibly empty) subset of {extension, restriction}
type derivationSet string

// {extension, restriction, list, union}
type typeDerivationControl string

// {#all} or (possibly empty) subset of {extension, restriction, list, union}
type fullDerivationSet string

type XMLSchema struct {
	// Composition
	Include    []XMLInclude    `xml:"include"`
	Import     []XMLImport     `xml:"import"`
	Redefine   []XMLRedefine   `xml:"redefine"`
	Override   []XMLOverride   `xml:"override"`
	Annotation []XMLAnnotation `xml:"annotation"`

	DefaultOpenContent *XMLDefaultOpenContent `xml:"defaultOpenContent"`
	Annotation         []XMLAnnotation        `xml:"annotation"`
	XMLSchemaTops
}

type XMLAttribute struct {
	Default         string `xml:"default,attr"`
	Fixed           string `xml:"fixed,attr"`
	Form            string `xml:"form,attr"`
	Id              string `xml:"id,attr"`
	Name            NCName `xml:"name,attr"`
	Ref             QName  `xml:"ref,attr"`
	TargetNamespace anyURI `xml:"targetNamespace,attr"`
	Type            QName  `xml:"type,attr"`
	Use             string `xml:"use,attr"`
	Inheritable     bool   `xml:"inheritable,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
	SimpleType *XMLSimpleType `xml:"simpleType"`
}

type XMLAnnotation struct {
	Id string `xml:"id,attr"`

	AppInfo       *XMLAppInfo       `xml:"appinfo"`
	Documentation *XMLDocumentation `xml:"documentation"`
}

type XMLAppInfo struct {
	Source  anyURI `xml:"source,attr"`
	Content string `xml:",chardata"`
}

type XMLDocumentation struct {
	Source  anyURI `xml:"source,attr"`
	Content string `xml:",chardata"`
}

type XMLMinExclusive struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLMinInclusive struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLMaxExclusive struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLMaxInclusive struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLTotalDigits struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value int    `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLFractionDigits struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value int    `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLLength struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value int    `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLMinLength struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value int    `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLMaxLength struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value int    `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLEnumeration struct {
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLWhiteSpace struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	// (collapse | preserve | replace)
	Value string `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLPattern struct {
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLAssertion struct {
	Id                    string `xml:"id,attr"`
	Test                  string `xml:"test,attr"`
	XPathDefaultNamespace string `xml:"xpathDefaultNamespace,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLExplicitTimezone struct {
	Fixed string `xml:"fixed,attr"`
	Value NCName `xml:"value,attr"`

	Annotation *XMLAnnotation `xml:"annotation"`
}

type XMLLocalSimpleType struct {
	XMLSimpleType
}

type XMLTopLevelSimpleType struct {
	XMLSimpleType
	Final string `xml:"final,attr"`
	Name  NCName `xml:"name,attr"`
}

type XMLSimpleRestrictionModel struct {
	// Annotated
	Annotation *XMLAnnotation `xml:"annotation"`
	Id         string         `xml:"id,attr"`

	SimpleType *XMLLocalSimpleType `xml:"simpleType"`

	MinExclusive     []XMLMinExclusive     `xml:"minExclusive"`
	MinInclusive     []XMLMinInclusive     `xml:"minInclusive"`
	MaxExclusive     []XMLMaxExclusive     `xml:"maxExclusive"`
	MaxInclusive     []XMLMaxInclusive     `xml:"maxInclusive"`
	TotalDigits      []XMLTotalDigits      `xml:"totalDigits"`
	FractionDigits   []XMLFractionDigits   `xml:"fractionDigits"`
	Length           []XMLLength           `xml:"length"`
	MinLength        []XMLMinLength        `xml:"minLength"`
	MaxLength        []XMLMaxLength        `xml:"maxLength"`
	Enumeration      []XMLEnumeration      `xml:"enumeration"`
	WhiteSpace       []XMLWhiteSpace       `xml:"whiteSpace"`
	Pattern          []XMLPattern          `xml:"pattern"`
	Assertion        []XMLAssertion        `xml:"assertion"`
	ExplicitTimezone []XMLExplicitTimezone `xml:"explicitTimezone"`
}

type XMLSimpleType struct {
	Id string `xml:"id,attr"`

	Restriction *struct {
		Base QName  `xml:"base,attr"`
		Id   string `xml:"id,attr"`

		Annotation *XMLAnnotation `xml:"annotation"`
	} `xml:"restriction"`
	List *struct {
		Id       string `xml:"id,attr"`
		ItemType QName  `xml:"itemType,attr"`

		Annotation *XMLAnnotation `xml:"annotation"`
		SimpleType *XMLSimpleType `xml:"simpleType"`
	} `xml:"list"`
	Union *struct {
		Id          string      `xml:"id,attr"`
		MemberTypes ListOfQName `xml:"memberTypes,attr"`

		Annotation  *XMLAnnotation   `xml:"annotation"`
		SimpleTypes []*XMLSimpleType `xml:"simpleType"`
	} `xml:"union"`
}

type XMLElement struct {
	Abstract bool `xml:"abstract,attr"`
	// (#all | List of (extension | restriction | substitution))
	Block     string `xml:"block,attr"`
	Default   string `xml:"default,attr"`
	Final     string `xml:"final,attr"`
	Fixed     string `xml:"fixed,attr"`
	Form      string `xml:"form,attr"`
	Id        string `xml:"id,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`
	MinOccurs int    `xml:"minOccurs,attr"`
}

type XMLDefaultOpenContent struct {
}

type XMLInclude struct {
}

type XMLImport struct {
}

type XMLRedefine struct {
}

type XMLOverride struct {
}

type XMLComplexType struct {
}

type XMLGroup struct {
}

type XMLAttributeGroup struct {
}

type XMLNotation struct {
}
