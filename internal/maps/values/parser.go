package values

import "github.com/ldez/structtags/parser"

// Parse parses a struct tag to a `map[string][]string`.
// The value is split on comma.
func Parse(tag string, escapeComma bool) (Tag, error) {
	return parser.Tag(tag, NewFiller(escapeComma))
}
