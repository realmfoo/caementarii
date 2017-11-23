package simple04

import (
	"encoding/xml"
)

type Employee struct {
	XMLName   xml.Name `xml:"urn:caementarii:simple employee"`
	Firstname string   `xml:"firstname"`
	Lastname  string   `xml:"lastname"`
	Address   string   `xml:"address"`
	City      string   `xml:"city"`
	Country   string   `xml:"country"`
}
