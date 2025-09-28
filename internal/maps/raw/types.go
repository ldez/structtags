package raw

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
)

// config for the parser.
type config struct {
	// DuplicateKeysMode allows duplicate keys.
	DuplicateKeysMode DuplicateKeysMode
}

type Option func(*config)

func WithDuplicateKeysMode(mode DuplicateKeysMode) Option {
	return func(options *config) {
		options.DuplicateKeysMode = mode
	}
}

// Tag is a key/value map.
type Tag map[string]string

func (m Tag) String() string {
	var b strings.Builder

	for k, v := range m {
		b.WriteString(fmt.Sprintf("%s:%q ", k, v))
	}

	return strings.TrimSuffix(b.String(), " ")
}
