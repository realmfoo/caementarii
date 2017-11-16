package simple03

import (
	"encoding/xml"
)

type TestOptional struct {
	XMLName xml.Name `xml:"urn:caementarii:simple testOptional"`
	Age     *string  `xml:"age,attr,omitempty"`
}

type TestRequired struct {
	XMLName xml.Name `xml:"urn:caementarii:simple testRequired"`
	Age     string   `xml:"age,attr"`
}
