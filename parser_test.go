package spam

import (
	"fmt"
	"strings"
	"testing"
)

/*
func TestParser(t *testing.T) {
	src := `hello, {harrison|barty}. How are you{!|?}`

	p := NewParser(NewLexer(strings.NewReader(src)))

	spam := p.Parse()
	fmt.Println(spam.String())

	fmt.Println(p.buf)
}*/

func TestNestedParser(t *testing.T) {
	src := `... {hi{harrison|barry|something}|this is kill{hey|party|something}}`

	p := NewParser(NewLexer(strings.NewReader(src)))

	spam := p.Parse()
	fmt.Println(spam.String())

	fmt.Println(p.buf)
}
