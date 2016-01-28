package spam

import (
	"bytes"
	"fmt"
	"math/rand"
)

type Spam struct {
	words []fmt.Stringer
}

func (s Spam) String() string {
	var buf bytes.Buffer
	for _, word := range s.words {
		buf.WriteString(word.String())
	}

	return buf.String()
}

type Text struct {
	Content string
}

func (t Text) String() string {
	return t.Content
}

type Option struct {
	Either []string
}

func (o Option) String() string {
	word := o.Either[rand.Intn(len(o.Either))]

	return word
}
