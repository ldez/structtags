package multikeys

type Tag map[string][]string

type Filler struct {
	data Tag
}

func (f *Filler) Data() Tag {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	if f.data == nil {
		f.data = Tag{}
	}

	f.data[key] = append(f.data[key], value)

	return nil
}
