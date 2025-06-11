package simple03

import (
	"bytes"
	"encoding/xml"
	"github.com/realmfoo/caementarii"
	"github.com/realmfoo/caementarii/xsd"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSimple03(t *testing.T) {
	data, err := os.ReadFile("simple03.xsd")
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
		PkgName: "simple03",
	}
	err = g.Generate(&s, buf)
	if err != nil {
		t.Fatal(err)
	}

	expected, _ := os.ReadFile("simple03.go")
	assert.Equal(t, string(expected), buf.String())
}
