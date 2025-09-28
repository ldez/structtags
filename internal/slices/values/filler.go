package values

import (
	"github.com/ldez/structtags/parser"
)

type Filler struct {
	data        Tags
	escapeComma bool

	keys map[string]struct{}
}

func NewFiller(escapeComma bool) *Filler {
	return &Filler{
		escapeComma: escapeComma,
		keys:        make(map[string]struct{}),
	}
}

func (f *Filler) Data() Tags {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	if _, ok := f.keys[key]; ok {
		// Ignore duplicated keys.
		// TODO(ldez) add an option to through an error.
		return nil
	}

	f.keys[key] = struct{}{}

	values, err := parser.Value(value, f.escapeComma)
	if err != nil {
		return err
	}

	f.data = append(f.data, Tag{
		Key:    key,
		Values: values,
	})

	return nil
}
