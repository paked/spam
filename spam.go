package spam

import "io"

func New(r io.Reader) *Spam {
	l := NewLexer(r)
	p := NewParser(l)

	return p.Parse()
}
