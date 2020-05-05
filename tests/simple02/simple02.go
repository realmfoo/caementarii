package simple02

import (
	"encoding/xml"
)

type Person struct {
	XMLName     xml.Name `xml:"urn:caementarii:simple person"`
	RequiredAge string   `xml:"requiredAge,attr"`
	Age         *string  `xml:"age,attr,omitempty"`
	Disabled    *bool    `xml:"disabled,attr,omitempty"`
}

type PersonName struct {
	XMLName  xml.Name `xml:"urn:caementarii:simple personName"`
	Title    *string  `xml:"title"`
	Forename []string `xml:"forename"`
	Surname  string   `xml:"surname"`
}
