package simple01

import (
	"encoding/xml"
)

type Lastname struct {
	XMLName xml.Name `xml:"urn:caementarii:simple lastname"`
	Value   string   `xml:",chardata"`
}
