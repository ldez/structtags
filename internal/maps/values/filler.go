package values

import "github.com/ldez/structtags/parser"

type Filler struct {
	data        Tag
	escapeComma bool
}

func NewFiller(escapeComma bool) *Filler {
	return &Filler{
		escapeComma: escapeComma,
	}
}

func (f *Filler) Data() Tag {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	if f.data == nil {
		f.data = Tag{}
	}

	values, err := parser.Value(value, f.escapeComma)
	if err != nil {
		return err
	}

	f.data[key] = append(f.data[key], values...)

	return nil
}
