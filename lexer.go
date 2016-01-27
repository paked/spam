package spam

import (
	"bufio"
	"bytes"
	"io"
)

type Lexer struct {
	r *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		r: bufio.NewReader(r),
	}
}

func (l *Lexer) Scan() (Token, string) {
	ch := l.read()

	if l.isWhitespace(ch) {
		l.unread()

		return l.scanWhitespace()
	}

	switch ch {
	case '{':
		return OpenStatement, string(ch)
	case '}':
		return CloseStatement, string(ch)
	case '|':
		return Or, string(ch)
	case eof:
		return EOF, ""
	}

	return Illegal, string(ch)
}

func (l *Lexer) isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (l *Lexer) scanWhitespace() (Token, string) {
	var buf bytes.Buffer

	for {
		ch := l.read()

		if ch == eof {
			break
		} else if !l.isWhitespace(ch) {
			l.unread()
			break
		}

		buf.WriteRune(ch)
	}

	return Whitespace, buf.String()
}

func (l *Lexer) read() rune {
	ch, _, err := l.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

func (l *Lexer) unread() {
	l.r.UnreadRune()
}

var eof = rune(0)
