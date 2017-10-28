package xmlschema

import (
	"encoding/xml"
	"fmt"
	"github.com/realmfoo/caementarii/xsd"
	"io/ioutil"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	data, err := ioutil.ReadFile("XMLSchema.xsd")
	if err != nil {
		t.Fatal(err)
	}

	schema := xsd.Schema{}
	err = xml.Unmarshal(data, &schema)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", schema)
}
