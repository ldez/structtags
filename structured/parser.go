package structured

import "github.com/ldez/structtags/parser"

// Parse parses a struct tag to a [Tag].
// Allows modifying the struct tags.
// The value is split on comma.
func Parse(tag string, opt *Options) (*Tag, error) {
	var escapeComma bool

	var allowDuplicateKeys bool

	if opt != nil {
		escapeComma = opt.EscapeComma
		allowDuplicateKeys = opt.AllowDuplicateKeys
	}

	if tag == "" {
		return NewTag(escapeComma, allowDuplicateKeys), nil
	}

	return parser.Tag(tag, NewFiller(escapeComma, allowDuplicateKeys))
}
