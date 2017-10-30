package simple01

import (
	"bytes"
	"encoding/xml"
	"github.com/realmfoo/caementarii"
	"github.com/realmfoo/caementarii/xsd"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSimple01(t *testing.T) {
	data, err := ioutil.ReadFile("simple01.xsd")
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

	expected, _ := ioutil.ReadFile("simple01.go")
	assert.Equal(t, string(expected), buf.String())
}

func TestLastnameMarshaler(t *testing.T) {
	n := Lastname{Value: "Some value"}
	data, e := xml.Marshal(n)
	if e != nil {
		t.Fatal(e)
	}

	assert.Equal(t, `<lastname xmlns="urn:caementarii:simple">Some value</lastname>`, string(data))
}

func TestLastnameUnmarshaler(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{`<lastname xmlns="urn:caementarii:simple">Some value</lastname>`, `Some value`},
		{`<lastname xmlns="urn:caementarii:simple"/>`, ``},
	}

	for _, tt := range tests {
		n := Lastname{}
		e := xml.Unmarshal([]byte(tt.in), &n)
		if e != nil {
			t.Fatal(e)
		}

		assert.Equal(t, tt.out, n.Value)
	}
}
