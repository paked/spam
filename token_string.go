// generated by stringer -type=Token; DO NOT EDIT

package spam

import "fmt"

const _Token_name = "IllegalEOFWhitespaceAnyOpenStatementCloseStatementOr"

var _Token_index = [...]uint8{0, 7, 10, 20, 23, 36, 50, 52}

func (i Token) String() string {
	i -= 1
	if i < 0 || i >= Token(len(_Token_index)-1) {
		return fmt.Sprintf("Token(%d)", i+1)
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}
