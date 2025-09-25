package fatihext

import (
	"github.com/fatih/structtag"
	"github.com/ldez/structtags/parser"
)

type Filler struct {
	data []*structtag.Tag
}

func (f *Filler) Data() []*structtag.Tag {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	values, err := parser.Value(value)
	if err != nil {
		return err
	}

	name := values[0]

	options := values[1:]
	if len(options) == 0 {
		options = nil
	}

	f.data = append(f.data, &structtag.Tag{
		Key:     key,
		Name:    name,
		Options: options,
	})

	return nil
}
