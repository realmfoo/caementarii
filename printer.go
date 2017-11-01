package goxsd

import "io"

type ctrlSymbol int

const (
	none ctrlSymbol = iota
	semi
	blank
	newline
	indent
	outdent
	// comment
	// eolComment
)

type printer struct {
	output io.Writer

	indent  int // current indentation level
	nlcount int // number of consecutive newlines

	pending []whitespace // pending whitespace
	lastTok token        // last token (after any pending semi) processed by print
}

type whitespace struct {
	last token
	kind ctrlSymbol
	//text string // comment text (possibly ""); valid if kind == comment
}

func (p *printer) print(args ...interface{}) {
	for i := 0; i < len(args); i++ {
		switch x := args[i].(type) {
		case nil:
		case Node:
			p.printNode(x)

		case token:
			var s string
			if x == _Name {
				i++
				if i >= len(args) {
					panic("missing string argument after _Name")
				}
				s = args[i].(string)
			} else {
				s = x.String()
			}

			if x == _Semi {
				// delay printing of semi
				p.addWhitespace(semi, "")
			} else {
				p.flush(x)
				p.writeString(s)
				p.nlcount = 0
				p.lastTok = x
			}

		case ctrlSymbol:
			switch x {
			case none, semi /*, comment*/ :
				panic("unreachable")
			}
			p.addWhitespace(x, "")
		}
	}
}

func (p *printer) addWhitespace(kind ctrlSymbol, text string) {
	p.pending = append(p.pending, whitespace{p.lastTok, kind /*text*/})
	switch kind {
	case semi:
		p.lastTok = _Semi
	case newline:
		p.lastTok = 0
		// TODO(gri) do we need to handle /*-style comments containing newlines here?
	}
}

func (p *printer) printNode(n Node) {
	switch n := n.(type) {
	case nil:
	case *Name:
		p.print(_Name, n.Value) // _Name requires actual value following immediately

	case *TypeDecl:
		if n.Group == nil {
			p.print(_Type, blank)
		}
		p.print(n.Name, blank)
		if n.Alias {
			p.print(_Assign, blank)
		}
		p.print(n.Type)

	}
}

func (p *printer) write(data []byte) {
	_, err := p.output.Write(data)
	if err != nil {
		panic(err)
	}
}

var (
	tabBytes    = []byte("\t\t\t\t\t\t\t\t")
	newlineByte = []byte("\n")
	blankByte   = []byte(" ")
)

func (p *printer) writeBytes(data []byte) {
	if len(data) == 0 {
		panic("expected non-empty []byte")
	}
	p.write(data)
}

func (p *printer) writeString(s string) {
	p.writeBytes([]byte(s))
}

func (p *printer) flush(next token) {
	// eliminate semis and redundant whitespace
	sawNewline := next == _EOF
	sawParen := next == _Rparen || next == _Rbrace
	for i := len(p.pending) - 1; i >= 0; i-- {
		switch p.pending[i].kind {
		case semi:
			k := semi
			if sawParen {
				sawParen = false
				k = none // eliminate semi
			} else if sawNewline && impliesSemi(p.pending[i].last) {
				sawNewline = false
				k = none // eliminate semi
			}
			p.pending[i].kind = k
		case newline:
			sawNewline = true
		case blank, indent, outdent:
			// nothing to do
			// case comment:
			// 	// A multi-line comment acts like a newline; and a ""
			// 	// comment implies by definition at least one newline.
			// 	if text := p.pending[i].text; strings.HasPrefix(text, "/*") && strings.ContainsRune(text, '\n') {
			// 		sawNewline = true
			// 	}
			// case eolComment:
			// 	// TODO(gri) act depending on sawNewline
		default:
			panic("unreachable")
		}
	}

	// print pending
	prev := none
	for i := range p.pending {
		switch p.pending[i].kind {
		case none:
			// nothing to do
		case semi:
			p.writeString(";")
			p.nlcount = 0
			prev = semi
		case blank:
			if prev != blank {
				// at most one blank
				p.writeBytes(blankByte)
				p.nlcount = 0
				prev = blank
			}
		case newline:
			const maxEmptyLines = 1
			if p.nlcount <= maxEmptyLines {
				p.write(newlineByte)
				p.nlcount++
				prev = newline
			}
		case indent:
			p.indent++
		case outdent:
			p.indent--
			if p.indent < 0 {
				panic("negative indentation")
			}
			// case comment:
			// 	if text := p.pending[i].text; text != "" {
			// 		p.writeString(text)
			// 		p.nlcount = 0
			// 		prev = comment
			// 	}
			// 	// TODO(gri) should check that line comments are always followed by newline
		default:
			panic("unreachable")
		}
	}

	p.pending = p.pending[:0] // re-use underlying array
}

func impliesSemi(tok token) bool {
	switch tok {
	case _Name,
		_Break, _Continue, _Fallthrough, _Return,
		/*_Inc, _Dec,*/ _Rparen, _Rbrack, _Rbrace: // TODO(gri) fix this
		return true
	}
	return false
}
