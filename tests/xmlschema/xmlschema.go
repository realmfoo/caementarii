package xmlschema

import (
	"encoding/xml"
)

type Import struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema import"`
}

type Union struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema union"`
}

type MinInclusive struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema minInclusive"`
}

type Assertion struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema assertion"`
}

type AnyAttribute struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema anyAttribute"`
}

type Redefine struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema redefine"`
}

type Notation struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema notation"`
}

type Restriction struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema restriction"`
}

type MaxExclusive struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema maxExclusive"`
}

type MaxInclusive struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema maxInclusive"`
}

type TotalDigits struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema totalDigits"`
}

type AttributeGroup struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema attributeGroup"`
}

type ComplexContent struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema complexContent"`
}

type Choice struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema choice"`
}

type Field struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema field"`
}

type Pattern struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema pattern"`
}

type Schema struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema schema"`
}

type Sequence struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema sequence"`
}

type Include struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema include"`
}

type Key struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema key"`
}

type Appinfo struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema appinfo"`
}

type Annotation struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema annotation"`
}

type XML struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema XML"`
}

type Enumeration struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema enumeration"`
}

type ComplexType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema complexType"`
}

type Facet struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema facet"`
}

type MinExclusive struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema minExclusive"`
}

type MinLength struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema minLength"`
}

type ExplicitTimezone struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema explicitTimezone"`
}

type Override struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema override"`
}

type Group struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema group"`
}

type Attribute struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema attribute"`
}

type SimpleType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
}

type Length struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema length"`
}

type All struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema all"`
}

type Documentation struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema documentation"`
}

type XMLSchemaStructures struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema XMLSchemaStructures"`
}

type MaxLength struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema maxLength"`
}

type SimpleContent struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema simpleContent"`
}

type DefaultOpenContent struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema defaultOpenContent"`
}

type Element struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema element"`
}

type Any struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema any"`
}

type Selector struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema selector"`
}

type Unique struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema unique"`
}

type Keyref struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema keyref"`
}

type List struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema list"`
}

type OpenContent struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema openContent"`
}

type WhiteSpace struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema whiteSpace"`
}

type FractionDigits struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema fractionDigits"`
}
