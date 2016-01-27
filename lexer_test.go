package spam

import (
	"fmt"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	src := `hello { a | b }`

	l := NewLexer(strings.NewReader(src))

	for {
		tok, val := l.Scan()

		if tok == EOF {
			break
		}

		fmt.Println(tok, val)
	}
}
