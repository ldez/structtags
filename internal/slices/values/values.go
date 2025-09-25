package values

import (
	"github.com/ldez/structtags/parser"
)

type Tag struct {
	Key    string
	Values []string
}

type Filler struct {
	data        []Tag
	escapeComma bool
}

func NewFiller(escapeComma bool) *Filler {
	return &Filler{
		escapeComma: escapeComma,
	}
}

func (f *Filler) Data() []Tag {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
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
