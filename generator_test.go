package goxsd

import (
	"encoding/xml"
	"github.com/realmfoo/caementarii/xsd"
	"io/ioutil"
	"testing"
)

func TestGenerator(t *testing.T) {
	data, err := ioutil.ReadFile("tests/XMLSchema.xsd")
	if err != nil {
		t.Fatal(err)
	}

	s := xsd.Schema{}
	err = xml.Unmarshal(data, &s)
	if err != nil {
		t.Fatal(err)
	}

	g := Generator{
		PkgName: "xmlschema",
	}
	g.Generate(&s)
}
