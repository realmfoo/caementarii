package goxsd

import (
	"bufio"
	"encoding/xml"
	"fmt"
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
		ImportResolver: func(namespace string, schemaLocation string) (*xsd.Schema, error) {
			if namespace != nsXml {
				return nil, fmt.Errorf("could not find a location of %s", namespace)
			}

			data, err := ioutil.ReadFile("tests/xmlschema/xml.xsd")
			if err != nil {
				return nil, err
			}

			s := xsd.Schema{}
			err = xml.Unmarshal(data, &s)
			if err != nil {
				return nil, err
			}
			return &s, nil
		},
	}
	err = g.Generate(&s, w)
	if err != nil {
		t.Fatal(err)
	}
}
