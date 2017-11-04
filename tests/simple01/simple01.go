package simple01

import (
	"encoding/xml"
)

var nsLastnameQName = xml.Name{Space: "urn:caementarii:simple", Local: "lastname"}

type Lastname string

func (t *Lastname) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return d.DecodeElement((*string)(t), &start)
}

func (t Lastname) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = nsLastnameQName
	return e.EncodeElement(string(t), start)
}
