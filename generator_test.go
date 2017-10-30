package goxsd

import (
	"bufio"
	"encoding/xml"
	"github.com/realmfoo/caementarii/xsd"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerator(t *testing.T) {
	data, err := ioutil.ReadFile("tests/xmlschema/XMLSchema.xsd")
	if err != nil {
		t.Fatal(err)
	}

	s := xsd.Schema{}
	err = xml.Unmarshal(data, &s)
	if err != nil {
		t.Fatal(err)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	g := Generator{
		PkgName: "xmlschema",
	}
	g.Generate(&s, w)
}
