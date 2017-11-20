package simple04

import (
	"encoding/xml"
)

type PersoninfoType struct {
	Firstname string `xml:"firstname"`
	Lastname  string `xml:"lastname"`
}

type FullpersoninfoType struct {
	PersoninfoType
	Address string `xml:"address"`
	City    string `xml:"city"`
	Country string `xml:"country"`
}

type Employee struct {
	XMLName xml.Name `xml:"urn:caementarii:simple employee"`
	FullpersoninfoType
}
