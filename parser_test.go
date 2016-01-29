package spam

import (
	"fmt"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	src := `hello, {harrison|barty}. How are you{!|?}`

	p := NewParser(NewLexer(strings.NewReader(src)))

	spam := p.Parse()
	fmt.Println(spam.String())

	fmt.Println(p.buf)
}
