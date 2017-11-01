package goxsd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeDecl(t *testing.T) {
	d := &TypeDecl{Name: &Name{Value: "Lastname"}, Type: &Name{Value: "string"}}
	buf := new(bytes.Buffer)
	p := printer{output: buf}
	p.print(d)
	p.flush(_EOF)
	assert.Equal(t, "type Lastname string", buf.String())
}
