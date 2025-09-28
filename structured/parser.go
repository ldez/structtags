package structured

import "github.com/ldez/structtags/parser"

// Parse parses a struct tag to a [Tag].
// Allows modifying the struct tags.
// The value is split on comma.
func Parse(tag string, options *Options) (*Tag, error) {
	var escapeComma bool

	var duplicateKeysMode DuplicateKeysMode

	if options != nil {
		escapeComma = options.EscapeComma
		duplicateKeysMode = options.DuplicateKeysMode
	}

	if tag == "" {
		return NewTag(escapeComma, duplicateKeysMode), nil
	}

	return parser.Tag(tag, NewFiller(escapeComma, duplicateKeysMode))
}
