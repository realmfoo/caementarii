package simple01

import (
	"bytes"
	"encoding/xml"
	"github.com/realmfoo/caementarii"
	"github.com/realmfoo/caementarii/xsd"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSimple01(t *testing.T) {
	data, err := os.ReadFile("simple01.xsd")
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
		PkgName: "simple01",
	}
	g.Generate(&s, buf)

	expected, _ := os.ReadFile("simple01.go")
	assert.Equal(t, string(expected), buf.String())
}

func TestLastnameMarshaler(t *testing.T) {
	lastname := Lastname("Some value")
	data, e := xml.Marshal(lastname)
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, `<lastname xmlns="urn:caementarii:simple">Some value</lastname>`, string(data))
}

func TestLastnameUnmarshaler(t *testing.T) {
	tests := []struct {
		in  string
		out Lastname
	}{
		{`<lastname xmlns="urn:caementarii:simple">Some value</lastname>`, `Some value`},
		{`<lastname xmlns="urn:caementarii:simple"/>`, ``},
	}

	for _, tt := range tests {
		var r Lastname
		e := xml.Unmarshal([]byte(tt.in), &r)
		if e != nil {
			t.Fatal(e)
		}

		assert.Equal(t, tt.out, r)
	}
}
