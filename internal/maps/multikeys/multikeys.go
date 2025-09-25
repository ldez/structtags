package multikeys

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

	f.data[key] = append(f.data[key], value)

	return nil
}
