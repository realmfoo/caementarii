package simple04

import (
	"bytes"
	"encoding/xml"
	"github.com/realmfoo/caementarii"
	"github.com/realmfoo/caementarii/xsd"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSimple04(t *testing.T) {
	data, err := os.ReadFile("simple04.xsd")
	if err != nil {
		t.Fatal(err)
	}

	s := xsd.Schema{}
	err = xml.Unmarshal(data, &s)
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)

	g := goxsd.Generator{
		PkgName: "simple04",
	}
	err = g.Generate(&s, buf)
	if err != nil {
		t.Fatal(err)
	}

	expected, _ := os.ReadFile("simple04.go")
	assert.Equal(t, string(expected), buf.String())
}

func TestMarshaler(t *testing.T) {
	employee := Employee{}
	employee.Firstname = "first"
	employee.Lastname = "last"
	employee.Address = "address"
	employee.City = "city"
	employee.Country = "country"

	data, e := xml.Marshal(employee)
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, `<employee xmlns="urn:caementarii:simple"><firstname>first</firstname><lastname>last</lastname><address>address</address><city>city</city><country>country</country></employee>`, string(data))
}

func TestUnmarshaler(t *testing.T) {
	in := `<employee xmlns="urn:caementarii:simple"><firstname>first</firstname><lastname>last</lastname><address>address</address><city>city</city><country>country</country></employee>`
	out := Employee{}

	e := xml.Unmarshal([]byte(in), &out)
	if e != nil {
		t.Fatal(e)
	}

	expected := Employee{XMLName: xml.Name{Space: "urn:caementarii:simple", Local: "employee"}}
	expected.Firstname = "first"
	expected.Lastname = "last"
	expected.Address = "address"
	expected.City = "city"
	expected.Country = "country"

	assert.Equal(t, expected, out)
}
