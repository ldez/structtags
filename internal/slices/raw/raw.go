package raw

type Tag struct {
	Key   string
	Value string
}

type Filler struct {
	data []Tag
}

func (f *Filler) Data() []Tag {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	f.data = append(f.data, Tag{
		Key:   key,
		Value: value,
	})

	return nil
}
