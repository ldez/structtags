package values

import (
	"fmt"
	"strings"
)

type DuplicateKeysMode int

const (
	// DuplicateKeysIgnore skips silently duplicate keys.
	DuplicateKeysIgnore DuplicateKeysMode = iota

	// DuplicateKeysDeny throws an error when duplicate keys are found.
	DuplicateKeysDeny

	// DuplicateKeysAllow NOT RECOMMENDED: this does not follow the struct tag conventions.
	DuplicateKeysAllow
)

// Options for the parser.
type Options struct {
	// EscapeComma is used to escape the comma character within the value.
	EscapeComma bool

	// DuplicateKeysMode allows duplicate keys.
	DuplicateKeysMode DuplicateKeysMode
}

// Tag is a key/values map.
type Tag map[string][]string

func (m Tag) String() string {
	var b strings.Builder

	for k, v := range m {
		b.WriteString(fmt.Sprintf("%s:%q ", k, strings.Join(v, ",")))
	}

	return strings.TrimSuffix(b.String(), " ")
}
