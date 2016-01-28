package spam

import (
	"bytes"
	"fmt"
)

type Parser struct {
	l *Lexer

	buf []buf
	n   int
}

func NewParser(l *Lexer) *Parser {
	return &Parser{
		l: l,
	}
}

func (p *Parser) Parse() *Spam {
	s := &Spam{}

	for {
		n := p.n
		tok, _ := p.scanSkipWhitespace()
		if tok == EOF {
			break
		}

		p.unscan()

		stmt, err := p.parseStatement()
		if err == nil {
			s.words = append(s.words, stmt)
		} else {
			p.reset(n)
		}

	}

	return s
}

func (p *Parser) parseStatement() (fmt.Stringer, error) {
	if p.is(OpenStatement) {
		return p.parseOptions()
	}

	return p.parseWords()
}

func (p *Parser) parseOptions() (fmt.Stringer, error) {
	tok, lit := p.scan()
	if tok != OpenStatement {
		return nil, fmt.Errorf("Syntax error: Was expecting '{', got '%v'", lit)
	}

	var firstWord bytes.Buffer

	for {
		tok, lit = p.scanSkipWhitespace()
		if tok == EOF {
			break
		}

		if tok != Illegal {
			p.unscan()
			break
		}

		firstWord.WriteString(lit)
	}

	tok, lit = p.scanSkipWhitespace()
	if tok != Or {
		return nil, fmt.Errorf("Syntax error: Was expecting '|', got '%v'", lit)
	}

	var secondWord bytes.Buffer

	for {
		tok, lit = p.scanSkipWhitespace()
		if tok == EOF {
			break
		}

		if tok != Illegal {
			p.unscan()
			break
		}

		secondWord.WriteString(lit)
	}

	tok, lit = p.scanSkipWhitespace()
	if tok != CloseStatement {
		return nil, fmt.Errorf("Syntax error: Was expecting '}', got '%v'", lit)
	}

	return Option{
		Either: []string{
			firstWord.String(),
			secondWord.String(),
		},
	}, nil
}

func (p *Parser) parseWords() (fmt.Stringer, error) {
	// white not EOF, or Or or CloseBracket add to lit buffer
	var buf bytes.Buffer
	for {
		tok, lit := p.scan()

		if tok == EOF {
			break
		}

		if tok == OpenStatement || tok == Or || tok == CloseStatement {
			p.unscan()
			break
		}

		buf.WriteString(lit)
	}

	return Text{buf.String()}, nil
}

func (p *Parser) is(ts ...Token) bool {
	for _, t := range ts {
		tok, _ := p.scan()

		defer func() { p.unscan() }()

		if t == Any {
			continue
		}

		if tok != t {
			return false
		}
	}

	return true
}

// If it can pull n from tokens, do that... else scan new tok
// and add it to the buf
func (p *Parser) scan() (Token, string) {
	defer func() {
		p.n++
	}()

	if p.n >= len(p.buf) {
		tok, lit := p.l.Scan()
		p.buf = append(p.buf, buf{
			tok: tok,
			lit: lit,
		})

		return tok, lit
	}

	b := p.buf[p.n]

	return b.tok, b.lit
}

func (p *Parser) unscan() {
	p.n--
}

func (p *Parser) scanSkipWhitespace() (Token, string) {
	for {
		tok, lit := p.scan()

		if tok != Whitespace {
			return tok, lit
		}
	}
}

func (p *Parser) reset(n int) {
	for p.n > n {
		p.unscan()
	}
}

type buf struct {
	tok Token
	lit string
}
