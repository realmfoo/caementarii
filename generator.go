package goxsd

import (
	"bufio"
	"github.com/realmfoo/caementarii/xsd"
	"os"
)

type Generator struct {
	pkgName string
}

func (g *Generator) generate(s *xsd.Schema) {

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	w.WriteString("package " + g.pkgName + "\n\n")
}
