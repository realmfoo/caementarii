package simple02

import (
	"encoding/xml"
)

type PersonName struct {
	XMLName  xml.Name `xml:"urn:caementarii:simple personName"`
	Title    *string  `xml:"title"`
	Forename []string `xml:"forename"`
	Surname  string   `xml:"surname"`
}

type Person struct {
	XMLName     xml.Name `xml:"urn:caementarii:simple person"`
	RequiredAge string   `xml:"requiredAge,attr"`
	Age         *string  `xml:"age,attr,omitempty"`
}
