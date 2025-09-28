package raw

import (
	"fmt"
	"strings"
)

type Tag map[string]string

func (m Tag) String() string {
	var b strings.Builder

	for k, v := range m {
		b.WriteString(fmt.Sprintf("%s:%q ", k, v))
	}

	return strings.TrimSuffix(b.String(), " ")
}

type Filler struct {
	data Tag
}

func (f *Filler) Data() Tag {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	if f.data != nil && f.data[key] != "" {
		// Ignore duplicated key.
		// TODO(ldez) add an option to through an error.
		return nil
	}

	if f.data == nil {
		f.data = Tag{}
	}

	f.data[key] = value

	return nil
}
