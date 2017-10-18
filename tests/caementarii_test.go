package tests

import (
	"encoding/xml"
	"fmt"
	"github.com/realmfoo/caementarii"
	"io/ioutil"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	data, err := ioutil.ReadFile("XMLSchema.xsd")
	if err != nil {
		t.Fatal(err)
	}

	schema := goxsd.XMLSchema{}
	err = xml.Unmarshal(data, &schema)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", schema)
}
