package values

import "github.com/ldez/structtags/parser"

// Parse parses a struct tag to a `map[string][]string`.
// The value is split on comma.
// Ignore duplicated keys.
func Parse(tag string, options *Options) (Tag, error) {
	var escapeComma bool

	var duplicateKeysMode DuplicateKeysMode

	if options != nil {
		escapeComma = options.EscapeComma
		duplicateKeysMode = options.DuplicateKeysMode
	}

	return parser.Tag(tag, NewFiller(escapeComma, duplicateKeysMode))
}
