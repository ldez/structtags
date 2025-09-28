package values

import "github.com/ldez/structtags/parser"

// Parse parses a struct tag to a slice of [Tag].
// The value is split on comma.
func Parse(tag string, escapeComma bool) (Tags, error) {
	return parser.Tag(tag, NewFiller(escapeComma))
}
