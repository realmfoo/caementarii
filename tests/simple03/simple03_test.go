package simple03

import (
	"bytes"
	"encoding/xml"
	"github.com/realmfoo/caementarii"
	"github.com/realmfoo/caementarii/xsd"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSimple03(t *testing.T) {
	data, err := ioutil.ReadFile("simple03.xsd")
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
	g.Generate(&s, buf)

	expected, _ := ioutil.ReadFile("simple03.go")
	assert.Equal(t, string(expected), buf.String())
}
