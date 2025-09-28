package raw

import "github.com/ldez/structtags/parser"

// Parse parses a struct tag to a `map[string]string`.
func Parse(tag string) (Tag, error) {
	return parser.Tag(tag, &Filler{})
}
