package simple01

import (
	"encoding/xml"
)

type Lastname string

var LastnameQName = xml.Name{Space: "urn:caementarii:simple", Local: "lastname"}

func (t *Lastname) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return d.DecodeElement((*string)(t), &start)
}

func (t Lastname) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = LastnameQName
	return e.EncodeElement(string(t), start)
}
