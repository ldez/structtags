package structured

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/ldez/structtags/parser"
)

// Options for the parser.
type Options struct {
	// EscapeComma is used to escape the comma character within the value.
	EscapeComma bool

	// AllowDuplicateKeys allows duplicate keys.
	// NOT RECOMMENDED TO USE IT: this does not follow the struct tag conventions.
	AllowDuplicateKeys bool
}

// Tag represents a struct tag.
type Tag struct {
	entries []*Entry

	escapeComma        bool
	allowDuplicateKeys bool
}

// NewTag creates a new [Tag].
func NewTag(escapeComma, allowDuplicateKeys bool) *Tag {
	return &Tag{
		escapeComma:        escapeComma,
		allowDuplicateKeys: allowDuplicateKeys,
	}
}

// Get returns the first entry with the given key.
func (t *Tag) Get(key string) *Entry {
	for _, tag := range t.entries {
		if tag != nil && tag.Key == key {
			return tag
		}
	}

	return nil
}

// GetAll returns all entries with the given key.
// NOT RECOMMENDED TO USE IT: this does not follow the struct tag conventions.
func (t *Tag) GetAll(key string) []*Entry {
	var entries []*Entry

	for _, tag := range t.entries {
		if tag != nil && tag.Key == key {
			entries = append(entries, tag)
		}
	}

	return entries
}

// Add adds a new entry to the [Tag].
func (t *Tag) Add(tag *Entry) error {
	if tag == nil {
		return nil
	}

	if !t.allowDuplicateKeys && t.Get(tag.Key) != nil {
		return fmt.Errorf("duplicate tag %q", tag.Key)
	}

	tag.escapeComma = t.escapeComma

	t.entries = append(t.entries, tag)

	return nil
}

// Delete deletes the entry with the given key.
func (t *Tag) Delete(key string) {
	t.entries = slices.DeleteFunc(t.entries, func(entry *Entry) bool {
		return entry != nil && entry.Key == key
	})
}

// Seq returns a sequence of entries.
func (t *Tag) Seq() iter.Seq[*Entry] {
	return func(yield func(*Entry) bool) {
		for _, entry := range t.entries {
			if entry == nil {
				continue
			}

			yield(entry)
		}
	}
}

// IsEmpty returns true if the [Tag] is empty.
func (t *Tag) IsEmpty() bool {
	return len(t.entries) == 0
}

// Sort sorts the entries alphabetically by key.
func (t *Tag) Sort() {
	slices.SortFunc(t.entries, func(a, b *Entry) int {
		return strings.Compare(a.Key, b.Key)
	})
}

// String returns the string representation of the [Tag].
func (t *Tag) String() string {
	var b strings.Builder

	for i, tag := range t.entries {
		b.WriteString(tag.String())

		if i != len(t.entries)-1 {
			b.WriteString(" ")
		}
	}

	return b.String()
}

// Entry represents a struct tag entry.
// An entry is composed of a key and a value.
type Entry struct {
	Key      string
	RawValue string

	escapeComma bool
}

// Values returns the values of the entry.
// When modifying the values, the result must be set [Entry.RawValue].
func (e *Entry) Values() (TagValues, error) {
	return parser.Value(e.RawValue, e.escapeComma)
}

// String returns the string representation of the entry.
func (e *Entry) String() string {
	return fmt.Sprintf("%s=%q", e.Key, e.RawValue)
}

// TagValues is a slice of values related to a key.
type TagValues []string

// Has checks if the values contain the given value.
func (t TagValues) Has(value string) bool {
	return slices.Contains(t, value)
}

// IsEmpty returns true if the values are empty.
func (t TagValues) IsEmpty() bool {
	return len(t) == 0 || len(t) == 1 && t[0] == ""
}

// String returns the string representation of the values.
func (t TagValues) String() string {
	return strings.Join(t, ",")
}
