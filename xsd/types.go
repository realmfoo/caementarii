package xsd

import (
	"encoding/xml"
)

// NCName represents XML "non-colonized" Names.
// white-space: collapse
type NCName string

// QName represents XML qualified names.
type QName string

// A list of QName
type ListOfQName []QName

// anyURI represents an Internationalized Resource Identifier Reference (IRI).
type anyURI string

// This group is for the elements which occur freely at the top level of schemas.
// All of their types are based on the "annotated" type by extension.
var schemaTopGroup = make(map[xml.Name]interface{})

func init() {
	schemaTopGroup[xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "simpleType"}] = SimpleType{}
	schemaTopGroup[xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "complexType"}] = ComplexType{}
	schemaTopGroup[xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "group"}] = Group{}
	schemaTopGroup[xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "attributeGroup"}] = AttributeGroup{}
	schemaTopGroup[xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "element"}] = Element{}
	schemaTopGroup[xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "attribute"}] = Attribute{}
	schemaTopGroup[xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "notation"}] = Notation{}
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

type Schema struct {
	Composition        []interface{}
	DefaultOpenContent *XMLDefaultOpenContent `xml:"defaultOpenContent"`
	Annotation         []Annotation           `xml:"annotation"`
	SchemaTop          []interface{}
}

//  <xs:choice>
//    <xs:element ref="xs:include"/>
//    <xs:element ref="xs:import"/>
//    <xs:element ref="xs:redefine"/>
//    <xs:element ref="xs:override"/>
//    <xs:element ref="xs:annotation"/>
//  </xs:choice>
func unmarshalCompositionGroupChoice(d *xml.Decoder, tok xml.Token) (interface{}, xml.Token, error) {
	var err error
	var r interface{}

	for {
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name {

			// <xs:element ref="xs:include"/>
			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "include"}:
				x := Include{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, err
				}
				r = x

			// <xs:element ref="xs:import"/>
			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "import"}:
				x := Import{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, err
				}
				r = x

			// <xs:element ref="xs:redefine"/>
			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "redefine"}:
				x := Redefine{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, err
				}
				r = x

			// <xs:element ref="xs:override"/>
			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "override"}:
				x := Override{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, err
				}
				r = x

			// <xs:element ref="xs:annotation"/>
			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "annotation"}:
				x := Annotation{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, err
				}
				r = x

			// unexpected element
			default:
				return nil, tok, nil
			}

			// read next token
			tok, err = d.Token()
			if err != nil {
				return r, tok, err
			}

			return r, tok, nil

		case xml.EndElement:
			return nil, tok, err
		}

		// read next token until xml.StartElement
		tok, err = d.Token()
		if err != nil {
			return r, tok, err
		}

	}

}

func (s *Schema) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	//s.Xmlns = make(map[string]string)
	//s.XMLName = start.Name
	//for _, attr := range start.Attr {
	//	if attr.Name.Space == "xmlns" {
	//		s.Xmlns[attr.Name.Local] = attr.Value
	//		continue
	//	}
	//
	//	switch attr.Name.Local {
	//	case "version":
	//		s.Version = attr.Value
	//	case "targetNamespace":
	//		s.TargetNamespace = attr.Value
	//	case "elementFormDefault":
	//		s.ElementFormDefault = attr.Value
	//	}
	//}
	tok, err := d.Token()
	if err != nil {
		return err
	}

	//  <xs:group name="composition" minOccurs="0" maxOccurs="unbounded">
	{
		for {
			//    <xs:choice>
			//      <xs:element ref="xs:include"/>
			//      <xs:element ref="xs:import"/>
			//      <xs:element ref="xs:redefine"/>
			//      <xs:element ref="xs:override"/>
			//      <xs:element ref="xs:annotation"/>
			//    </xs:choice>
			var x interface{}

			x, tok, err = unmarshalCompositionGroupChoice(d, tok)
			if err != nil {
				return err
			}

			// minOccurs="0"
			if x == nil {
				break
			}

			s.Composition = append(s.Composition, x)

			// maxOccurs="unbounded"
		}
	}

	// <xs:sequence minOccurs="0">
	//   <xs:element ref="xs:defaultOpenContent"/>
	//   <xs:element ref="xs:annotation" minOccurs="0" maxOccurs="unbounded"/>
	// </xs:sequence>

	// skip all other
Loop:
	for {
		switch tok.(type) {
		case xml.StartElement:
			d.Skip()
		case xml.EndElement:
			break Loop
		}

		tok, err = d.Token()
		if err != nil {
			return err
		}
	}

	return nil
}

type Attribute struct {
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

	Annotation *Annotation `xml:"annotation"`
	SimpleType *SimpleType `xml:"simpleType"`
}

type Annotation struct {
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

	Annotation *Annotation `xml:"annotation"`
}

type XMLMinInclusive struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLMaxExclusive struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLMaxInclusive struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLTotalDigits struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value int    `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLFractionDigits struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value int    `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLLength struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value int    `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLMinLength struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value int    `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLMaxLength struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	Value int    `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLEnumeration struct {
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLWhiteSpace struct {
	Fixed string `xml:"fixed,attr"`
	Id    string `xml:"id,attr"`
	// (collapse | preserve | replace)
	Value string `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLPattern struct {
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLAssertion struct {
	Id                    string `xml:"id,attr"`
	Test                  string `xml:"test,attr"`
	XPathDefaultNamespace string `xml:"xpathDefaultNamespace,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLExplicitTimezone struct {
	Fixed string `xml:"fixed,attr"`
	Value NCName `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type XMLLocalSimpleType struct {
	SimpleType
}

type XMLTopLevelSimpleType struct {
	SimpleType
	Final string `xml:"final,attr"`
	Name  NCName `xml:"name,attr"`
}

type XMLSimpleRestrictionModel struct {
	// Annotated
	Annotation *Annotation `xml:"annotation"`
	Id         string      `xml:"id,attr"`

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

type SimpleType struct {
	Id string `xml:"id,attr"`

	Restriction *struct {
		Base QName  `xml:"base,attr"`
		Id   string `xml:"id,attr"`

		Annotation *Annotation `xml:"annotation"`
	} `xml:"restriction"`
	List *struct {
		Id       string `xml:"id,attr"`
		ItemType QName  `xml:"itemType,attr"`

		Annotation *Annotation `xml:"annotation"`
		SimpleType *SimpleType `xml:"simpleType"`
	} `xml:"list"`
	Union *struct {
		Id          string      `xml:"id,attr"`
		MemberTypes ListOfQName `xml:"memberTypes,attr"`

		Annotation  *Annotation   `xml:"annotation"`
		SimpleTypes []*SimpleType `xml:"simpleType"`
	} `xml:"union"`
}

type Element struct {
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

type Include struct {
}

type Import struct {
	// Annotated
	Annotation *Annotation `xml:"annotation"`
	Id         string      `xml:"id,attr"`

	Namespace      anyURI `xml:"namespace,attr"`
	SchemaLocation anyURI `xml:"schemaLocation,attr"`
}

type Redefine struct {
}

type Override struct {
}

type ComplexType struct {
}

type Group struct {
}

type AttributeGroup struct {
}

type Notation struct {
}
