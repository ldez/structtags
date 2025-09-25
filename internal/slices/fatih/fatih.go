package fatih

import (
	"strings"

	"github.com/fatih/structtag"
)

type Filler struct {
	data []*structtag.Tag
}

func (f *Filler) Data() []*structtag.Tag {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	res := strings.Split(value, ",")

	name := res[0]

	options := res[1:]
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
