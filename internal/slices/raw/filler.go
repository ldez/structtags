package raw

type Filler struct {
	data Tags
}

func (f *Filler) Data() Tags {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	f.data = append(f.data, Tag{
		Key:   key,
		Value: value,
	})

	return nil
}
