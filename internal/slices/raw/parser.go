package raw

import "github.com/ldez/structtags/parser"

// Parse parses a struct tag to a slice of [Tag].
// Ignore duplicated keys.
func Parse(tag string, options *Options) (Tags, error) {
	var duplicateKeysMode DuplicateKeysMode

	if options != nil {
		duplicateKeysMode = options.DuplicateKeysMode
	}

	return parser.Tag(tag, NewFiller(duplicateKeysMode))
}
