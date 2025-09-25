package raw

import "fmt"

type Filler struct {
	data map[string]string
}

func (f *Filler) Data() map[string]string {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	if f.data != nil && f.data[key] != "" {
		return fmt.Errorf("duplicate tag %q", key)
	}

	if f.data == nil {
		f.data = map[string]string{}
	}

	f.data[key] = value

	return nil
}
