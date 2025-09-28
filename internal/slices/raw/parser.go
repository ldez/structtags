package raw

import "github.com/ldez/structtags/parser"

// Parse parses a struct tag to a slice of [Tag].
// Ignore duplicated keys.
func Parse(tag string) (Tags, error) {
	return parser.Tag(tag, NewFiller())
}
