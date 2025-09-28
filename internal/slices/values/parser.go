package values

import "github.com/ldez/structtags/parser"

// Parse parses a struct tag to a slice of [Tag].
// The value is split on comma.
// Ignore duplicated keys.
func Parse(tag string, options *Options) (Tags, error) {
	var escapeComma bool

	var duplicateKeysMode DuplicateKeysMode

	if options != nil {
		escapeComma = options.EscapeComma
		duplicateKeysMode = options.DuplicateKeysMode
	}

	return parser.Tag(tag, NewFiller(escapeComma, duplicateKeysMode))
}
