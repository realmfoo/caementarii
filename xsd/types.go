package xsd

import (
	"encoding/xml"
	"fmt"
	"math"
	"strconv"
)

var unbounded = math.MaxInt32

// NCName represents XML "non-colonized" Names.
// white-space: collapse
type NCName string

// QName represents XML qualified names.
type QName = string

// A list of QName
type ListOfQName []QName

// anyURI represents an Internationalized Resource Identifier Reference (IRI).
type anyURI = string

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
	XMLAttrs []xml.Attr `xml:"-"`

	AttributeFormDefault  string `xml:"attributeFormDefault,attr"`
	BlockDefault          string `xml:"blockDefault,attr"`
	DefaultAttributes     string `xml:"defaultAttributes,attr"`
	XpathDefaultNamespace string `xml:"xpathDefaultNamespace,attr"`
	ElementFormDefault    string `xml:"elementFormDefault,attr"`
	FinalDefault          string `xml:"finalDefault,attr"`
	Id                    string `xml:"id,attr"`
	TargetNamespace       string `xml:"targetNamespace,attr"`
	Version               string `xml:"version,attr"`
	//xml:lang = language

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

//<xs:group name="nestedParticle">
//  <xs:choice>
//    <xs:element name="element" type="xs:localElement"/>
//    <xs:element name="group" type="xs:groupRef"/>
//
//    <xs:element ref="xs:choice"/>
//    <xs:element ref="xs:sequence"/>
//    <xs:element ref="xs:any"/>
//  </xs:choice>
//</xs:group>
func unmarshalNestedParticleGroupChoice(d *xml.Decoder, tok xml.Token) (NestedParticle, xml.Token, error) {
	var err error
	var r NestedParticle

	for {
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name {

			// <xs:element name="element" type="xs:localElement"/>
			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "element"}:
				x := &Element{}
				if err = d.DecodeElement(x, &t); err != nil {
					return nil, tok, err
				}
				r = x

			// <xs:element name="group" type="xs:groupRef"/>
			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "group"}:
				x := &Group{}
				if err = d.DecodeElement(x, &t); err != nil {
					return nil, tok, err
				}
				r = x

			//    <xs:element ref="xs:choice"/>
			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "choice"}:
				x := &Choice{}
				if err = d.DecodeElement(x, &t); err != nil {
					return nil, tok, err
				}
				r = x

			//    <xs:element ref="xs:sequence"/>
			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "sequence"}:
				x := &Sequence{}
				if err = d.DecodeElement(x, &t); err != nil {
					return nil, tok, err
				}
				r = x

			//    <xs:element ref="xs:any"/>
			default:
				d.Skip()
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

func unmarshalSchemaTop(d *xml.Decoder, tok xml.Token) (interface{}, xml.Token, error) {
	var err error
	var r interface{}

	for {
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name {

			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "simpleType"}:
				x := SimpleType{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, fmt.Errorf("Failed to unmarshal <simpleType>: %s", err)
				}
				r = x

			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "complexType"}:
				x := ComplexType{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, fmt.Errorf("Failed to unmarshal <complexType>: %s", err)
				}
				r = x

			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "group"}:
				x := Group{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, fmt.Errorf("Failed to unmarshal <group>: %s", err)
				}
				r = x

			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "attributeGroup"}:
				x := AttributeGroup{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, fmt.Errorf("Failed to unmarshal <attributeGroup>: %s", err)
				}
				r = x

			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "element"}:
				x := Element{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, fmt.Errorf("Failed to unmarshal <element>: %s", err)
				}
				r = x

			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "attribute"}:
				x := Attribute{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, fmt.Errorf("Failed to unmarshal <attribute>: %s", err)
				}
				r = x

			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "notation"}:
				x := Notation{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, fmt.Errorf("Failed to unmarshal <notation>: %s", err)
				}
				r = x

			case xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "annotation"}:
				x := Annotation{}
				if err = d.DecodeElement(&x, &t); err != nil {
					return nil, tok, fmt.Errorf("Failed to unmarshal <annotation>: %s", err)
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
			return nil, tok, nil
		}

		// read next token until xml.StartElement
		tok, err = d.Token()
		if err != nil {
			return r, tok, err
		}

	}

}

func skipToStartElement(d *xml.Decoder, tok xml.Token) (xml.Token, error) {
	var err error
	for {
		switch tok.(type) {
		case xml.StartElement:
			return tok, nil
		case xml.EndElement:
			return nil, nil
		}

		tok, err = d.Token()
		if err != nil {
			return nil, err
		}
	}
}

func (s *Schema) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	//s.Xmlns = make(map[string]string)
	//s.XMLName = start.Name
	for _, attr := range start.Attr {
		switch attr.Name {
		case xml.Name{Space: "", Local: "version"}:
			s.Version = attr.Value
		case xml.Name{Space: "", Local: "targetNamespace"}:
			s.TargetNamespace = attr.Value
		case xml.Name{Space: "", Local: "elementFormDefault"}:
			s.ElementFormDefault = attr.Value
		case xml.Name{Space: "", Local: "blockDefault"}:
			s.BlockDefault = attr.Value
		case xml.Name{Space: "", Local: "defaultAttributes"}:
			s.DefaultAttributes = attr.Value
		case xml.Name{Space: "", Local: "xpathDefaultNamespace"}:
			s.XpathDefaultNamespace = attr.Value
		case xml.Name{Space: "", Local: "finalDefault"}:
			s.FinalDefault = attr.Value
		case xml.Name{Space: "", Local: "id"}:
			s.Id = attr.Value
		default:
			s.XMLAttrs = append(s.XMLAttrs, attr)
		}
	}

	tok, err := d.Token()
	if err != nil {
		return err
	}

	tok, err = skipToStartElement(d, tok)
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

	start = tok.(xml.StartElement)

	// <xs:sequence minOccurs="0">
	//   <xs:element ref="xs:defaultOpenContent"/>
	//   <xs:element ref="xs:annotation" minOccurs="0" maxOccurs="unbounded"/>
	// </xs:sequence>
	{
		if (start.Name == xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "defaultOpenContent"}) {
			if err = d.DecodeElement(&s.DefaultOpenContent, &start); err != nil {
				return err
			}

			tok, err = skipToStartElement(d, tok)
			if err != nil {
				return err
			}
			start = tok.(xml.StartElement)

			for {
				if (start.Name != xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "annotation"}) {
					break
				}

				x := Annotation{}
				if err = d.DecodeElement(&x, &start); err != nil {
					return err
				}
				s.Annotation = append(s.Annotation, x)

				tok, err = skipToStartElement(d, tok)
				if err != nil {
					return err
				}
				start = tok.(xml.StartElement)
			}
		}
	}

	// <xs:sequence minOccurs="0" maxOccurs="unbounded">
	//   <xs:choice>
	//	   <xs:element ref="xs:simpleType"/>
	//	   <xs:element ref="xs:complexType"/>
	//	   <xs:element ref="xs:group"/>
	//	   <xs:element ref="xs:attributeGroup"/>
	//	   <xs:element ref="xs:element"/>
	//	   <xs:element ref="xs:attribute"/>
	//	   <xs:element ref="xs:notation"/>
	//   </xs:choice>
	//   <xs:element ref="xs:annotation" minOccurs="0" maxOccurs="unbounded"/>
	// </xs:sequence>
	for {
		var x interface{}

		x, tok, err = unmarshalSchemaTop(d, tok)
		if err != nil {
			return err
		}

		// minOccurs="0"
		if x == nil {
			break
		}

		s.SchemaTop = append(s.SchemaTop, x)

		// maxOccurs="unbounded"
	}

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
	Default         *string `xml:"default,attr"`
	Fixed           *string `xml:"fixed,attr"`
	Form            string  `xml:"form,attr"`
	Id              string  `xml:"id,attr"`
	Name            string  `xml:"name,attr"`
	Ref             QName   `xml:"ref,attr"`
	TargetNamespace anyURI  `xml:"targetNamespace,attr"`
	Type            QName   `xml:"type,attr"`
	Use             string  `xml:"use,attr"`
	Inheritable     *bool   `xml:"inheritable,attr"`

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

type Pattern struct {
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type Assert struct {
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
	Pattern          []Pattern             `xml:"pattern"`
	Assertion        []Assert              `xml:"assertion"`
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
	Block             string  `xml:"block,attr"`
	Default           string  `xml:"default,attr"`
	Final             string  `xml:"final,attr"`
	Fixed             string  `xml:"fixed,attr"`
	Form              string  `xml:"form,attr"`
	Id                string  `xml:"id,attr"`
	MaxOccurs         *string `xml:"maxOccurs,attr"`
	MinOccurs         *int    `xml:"minOccurs,attr"`
	Name              string  `xml:"name,attr"`
	Nillable          bool    `xml:"nillable,attr"`
	Ref               string  `xml:"ref,attr"`
	SubstitutionGroup string  `xml:"substitutionGroup"`
	TargetNamespace   string  `xml:"targetNamespace,attr"`
	Type              QName   `xml:"type,attr"`

	Annotation  *Annotation  `xml:"annotation"`
	SimpleType  *SimpleType  `xml:"simpleType"`
	ComplexType *ComplexType `xml:"complexType"`
	Alternative *Alternative `xml:"alternative"`

	nestedParticle
}

type Alternative struct {
	Id                    string `xml:"id,attr"`
	Test                  string `xml:"test,attr"`
	Type                  string `xml:"type,attr"`
	XpathDefaultNamespace string `xml:"xpathDefaultNamespace"`

	Annotation  *Annotation  `xml:"annotation"`
	SimpleType  *SimpleType  `xml:"simpleType"`
	ComplexType *ComplexType `xml:"complexType"`
	Unique      []Unique     `xml:"unique"`
	Key         []Key        `xml:"key"`
	Keyref      []Keyref     `xml:"keyref"`
}

type Unique struct {
	Name string `xml:"name,attr"`
	Ref  string `xml:"ref,attr"`
}

type Key struct {
	Name string `xml:"name,attr"`
	Ref  string `xml:"ref,attr"`
}

type Keyref struct {
	Name  string `xml:"name,attr"`
	Ref   string `xml:"ref,attr"`
	Refer string `xml:"refer,attr"`
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
	Abstract               bool   `xml:"abstract,attr"`
	Block                  string `xml:"block,attr"`
	Final                  string `xml:"final,attr"`
	Id                     string `xml:"id,attr"`
	Mixed                  *bool  `xml:"mixed,attr"`
	Name                   string `xml:"name,attr"`
	DefaultAttributesApply bool   `xml:"defaultAttributesApply,attr"`

	Annotation      *Annotation      `xml:"annotation"`
	SimpleContent   *SimpleContent   `xml:"simpleContent"`
	ComplexContent  *ComplexContent  `xml:"complexContent"`
	Group           *Group           `xml:"group"`
	All             *All             `xml:"all"`
	Choice          *Choice          `xml:"choice"`
	Sequence        *Sequence        `xml:"sequence"`
	Attributes      []Attribute      `xml:"attribute"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	Assert          *Assert          `xml:"assert"`
}

type SimpleContent struct {
	Id string `xml:"id,attr"`

	Annotation  *Annotation `xml:"annotation"`
	Restriction *struct {
		Base QName  `xml:"base,attr"`
		Id   string `xml:"id,attr"`

		Annotation      *Annotation      `xml:"annotation"`
		SimpleType      *SimpleType      `xml:"simpleType"`
		FacetOrAny      []interface{}    `xml:"-"`
		Choice          *Choice          `xml:"choice"`
		Sequence        *Sequence        `xml:"sequence"`
		Attributes      []Attribute      `xml:"attribute"`
		AttributeGroups []AttributeGroup `xml:"attributeGroup"`
		AnyAtttribute   *AnyAttribute    `xml:"anyAttribute"`
		Assert          *Assert          `xml:"assert"`
	} `xml:"restriction"`
	Extension *struct {
		Base string `xml:"base,attr"`
		Id   string `xml:"id,attr"`

		Annotation      *Annotation      `xml:"annotation"`
		Attributes      []Attribute      `xml:"attribute"`
		AttributeGroups []AttributeGroup `xml:"attributeGroup"`
		AnyAtttribute   *AnyAttribute    `xml:"anyAttribute"`
		Assert          *Assert          `xml:"assert"`
	} `xml:"extension"`
}

type ComplexContent struct {
	Id    string `xml:"id,attr"`
	Mixed *bool  `xml:"mixed,attr"`

	Annotation  *Annotation `xml:"annotation"`
	Restriction *struct {
		Base string `xml:"base,attr"`
		Id   string `xml:"id,attr"`

		Annotation      *Annotation      `xml:"annotation"`
		Group           *Group           `xml:"group"`
		All             *All             `xml:"all"`
		Choice          *Choice          `xml:"choice"`
		Sequence        *Sequence        `xml:"sequence"`
		Attributes      []Attribute      `xml:"attribute"`
		AttributeGroups []AttributeGroup `xml:"attributeGroup"`
		Assert          *Assert          `xml:"assert"`
	} `xml:"restriction"`
	Extension *struct {
		Base string `xml:"base,attr"`
		Id   string `xml:"id,attr"`

		Annotation      *Annotation      `xml:"annotation"`
		Group           *Group           `xml:"group"`
		All             *All             `xml:"all"`
		Choice          *Choice          `xml:"choice"`
		Sequence        *Sequence        `xml:"sequence"`
		Attributes      []Attribute      `xml:"attribute"`
		AttributeGroups []AttributeGroup `xml:"attributeGroup"`
		Assert          *Assert          `xml:"assert"`
	} `xml:"extension"`
}

type AnyAttribute struct {
	Id               string `xml:"id,attr"`
	Namespace        string `xml:"namespace,attr"`
	NotNamespace     string `xml:"notNamespace,attr"`
	NotQName         string `xml:"notQName,attr"`
	ProcessConttents string `xml:"processContents,attr"`

	Annotation *Annotation `xml:"annotation"`
}

type AttributeGroup struct {
}

type Notation struct {
}

//-----------------------------------------------------------------------------
// typeDefParticle

type (
	Sequence struct {
		XMLAttrs []xml.Attr `xml:"-"`

		Id        string `xml:"id,attr"`
		MaxOccurs int    `xml:"maxOccurs,attr"`
		MinOccurs int    `xml:"minOccurs,attr"`

		Annotation *Annotation `xml:"annotation"`
		Content    []NestedParticle

		nestedParticle
	}

	Choice struct {
		nestedParticle
	}

	All struct {
	}

	Group struct {
		nestedParticle
	}
)

func (s *Sequence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	// Setup Defaults
	s.MinOccurs = 1
	s.MaxOccurs = 1

	//s.Xmlns = make(map[string]string)
	//s.XMLName = start.Name
	for _, attr := range start.Attr {
		switch attr.Name {
		case xml.Name{Space: "", Local: "id"}:
			s.Id = attr.Value
		case xml.Name{Space: "", Local: "maxOccurs"}:
			if attr.Value == "unbounded" {
				s.MaxOccurs = unbounded
			} else {
				s.MaxOccurs, err = strconv.Atoi(attr.Value)
				if err != nil {
					return fmt.Errorf("Invalid maxOccurs value: %v", err)
				}
			}
		case xml.Name{Space: "", Local: "minOccurs"}:
			s.MinOccurs, err = strconv.Atoi(attr.Value)
			if err != nil {
				return fmt.Errorf("Invalid minOccurs value: %v", err)
			}
		default:
			s.XMLAttrs = append(s.XMLAttrs, attr)
		}
	}

	tok, err := d.Token()
	if err != nil {
		return err
	}

	tok, err = skipToStartElement(d, tok)
	if err != nil {
		return err
	}

	// <xs:element ref="xs:annotation" minOccurs="0"/>
	for {
		if (start.Name != xml.Name{Space: "http://www.w3.org/2001/XMLSchema", Local: "annotation"}) {
			break
		}

		s.Annotation = &Annotation{}
		if err = d.DecodeElement(s.Annotation, &start); err != nil {
			return err
		}

		tok, err = skipToStartElement(d, tok)
		if err != nil {
			return err
		}
		start = tok.(xml.StartElement)
	}

	// <xs:group ref="xs:nestedParticle" minOccurs="0" maxOccurs="unbounded"/>
	{
		for {
			//<xs:choice>
			//  <xs:element name="element" type="xs:localElement"/>
			//  <xs:element name="group" type="xs:groupRef"/>
			//
			//  <xs:element ref="xs:choice"/>
			//  <xs:element ref="xs:sequence"/>
			//  <xs:element ref="xs:any"/>
			//</xs:choice>
			var x NestedParticle

			x, tok, err = unmarshalNestedParticleGroupChoice(d, tok)
			if err != nil {
				return err
			}

			// minOccurs="0"
			if x == nil {
				break
			}

			s.Content = append(s.Content, x)

			// maxOccurs="unbounded"
		}
	}

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

//-----------------------------------------------------------------------------
// Nested Particles

type NestedParticle interface {
	aNestedParticle()
}

type nestedParticle struct{}

func (*nestedParticle) aNestedParticle() {}
