package simple02

import (
	"bytes"
	"encoding/xml"
	"github.com/realmfoo/caementarii"
	"github.com/realmfoo/caementarii/xsd"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSimple02(t *testing.T) {
	data, err := ioutil.ReadFile("simple02.xsd")
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
		PkgName: "simple02",
	}
	g.Generate(&s, buf)

	expected, _ := ioutil.ReadFile("simple02.go")
	assert.Equal(t, string(expected), buf.String())
}

func TestMarshaler(t *testing.T) {
	tests := []struct {
		in  PersonName
		out string
	}{
		{PersonName{Surname: "Some value"}, `<personName xmlns="urn:caementarii:simple"><surname>Some value</surname></personName>`},
		{PersonName{Forename: []string{"a", "b"}, Surname: "Some value"}, `<personName xmlns="urn:caementarii:simple"><forename>a</forename><forename>b</forename><surname>Some value</surname></personName>`},
		{PersonName{}, `<personName xmlns="urn:caementarii:simple"><surname></surname></personName>`},
	}

	for _, tt := range tests {
		data, e := xml.Marshal(tt.in)
		if e != nil {
			t.Fatal(e)
		}
		assert.Equal(t, tt.out, string(data))
	}
}

func TestPersonMarshaler(t *testing.T) {
	tests := []struct {
		in  Person
		out string
	}{
		{Person{RequiredAge: "15", Age: xsstring("10")}, `<person xmlns="urn:caementarii:simple" requiredAge="15" age="10"></person>`},
		{Person{RequiredAge: "0"}, `<person xmlns="urn:caementarii:simple" requiredAge="0"></person>`},
		{Person{}, `<person xmlns="urn:caementarii:simple" requiredAge=""></person>`},
	}

	for _, tt := range tests {
		data, e := xml.Marshal(tt.in)
		if e != nil {
			t.Fatal(e)
		}
		assert.Equal(t, tt.out, string(data))

		var r Person
		e = xml.Unmarshal([]byte(tt.out), &r)
		if e != nil {
			t.Fatal(e)
		}

		tt.in.XMLName.Space = "urn:caementarii:simple"
		tt.in.XMLName.Local = "person"
		assert.Equal(t, tt.in, r)

	}
}

func TestUnmarshaler(t *testing.T) {
	ns := xml.Name{Space: "urn:caementarii:simple", Local: "personName"}
	tests := []struct {
		in  string
		out PersonName
	}{
		{`<personName xmlns="urn:caementarii:simple"><surname>Some value</surname></personName>`, PersonName{XMLName: ns, Surname: "Some value"}},
		{`<personName xmlns="urn:caementarii:simple"/>`, PersonName{XMLName: ns}},
	}

	for _, tt := range tests {
		var r PersonName
		e := xml.Unmarshal([]byte(tt.in), &r)
		if e != nil {
			t.Fatal(e)
		}

		assert.Equal(t, tt.out, r)
	}
}

func xsstring(s string) *string {
	return &s
}
