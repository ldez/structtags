package values

import (
	"fmt"
	"strings"

	"github.com/ldez/structtags/parser"
)

type Tags []Tag

func (t Tags) String() string {
	var b strings.Builder

	for _, e := range t {
		b.WriteString(fmt.Sprintf("%s:%q ", e.Key, strings.Join(e.Values, ",")))
	}

	return strings.TrimSuffix(b.String(), " ")
}

type Tag struct {
	Key    string
	Values []string
}

type Filler struct {
	data        Tags
	escapeComma bool
}

func NewFiller(escapeComma bool) *Filler {
	return &Filler{
		escapeComma: escapeComma,
	}
}

func (f *Filler) Data() Tags {
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
