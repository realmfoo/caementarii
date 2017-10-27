package tests

import (
	"encoding/xml"
	"github.com/realmfoo/caementarii"
	"github.com/realmfoo/caementarii/xsd"
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

	g := goxsd.Generator{
		PkgName: "xsd",
	}
	g.Generate(&s)

}
