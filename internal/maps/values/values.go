package values

import "github.com/ldez/structtags/parser"

type Filler struct {
	data map[string][]string
}

func (f *Filler) Data() map[string][]string {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	if f.data == nil {
		f.data = make(map[string][]string)
	}

	values, err := parser.Value(value)
	if err != nil {
		return err
	}

	f.data[key] = append(f.data[key], values...)

	return nil
}
