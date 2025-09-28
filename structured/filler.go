package structured

import (
	"fmt"
)

// Filler fills the tag from a struct tag.
type Filler struct {
	data *Tag

	escapeComma        bool
	allowDuplicateKeys bool
}

// NewFiller creates a new [Filler].
func NewFiller(escapeComma, allowDuplicateKeys bool) *Filler {
	return &Filler{
		escapeComma:        escapeComma,
		allowDuplicateKeys: allowDuplicateKeys,
	}
}

// Data returns the [Tag] filled by the struct tag content.
func (f *Filler) Data() *Tag {
	return f.data
}

// Fill fills the data from a struct tag.
func (f *Filler) Fill(key, value string) error {
	if f.data == nil {
		f.data = NewTag(f.escapeComma, f.allowDuplicateKeys)
	}

	if !f.allowDuplicateKeys && f.data.Get(key) != nil {
		return fmt.Errorf("duplicate tag %q", key)
	}

	return f.data.Add(&Entry{Key: key, RawValue: value})
}
