package spam

//go:generate stringer -type=Token
type Token int

const (
	Illegal Token = iota + 1
	EOF
	Whitespace
	Any

	// Brackets
	OpenStatement  // {
	CloseStatement // }

	// Operators
	Or // |
)
