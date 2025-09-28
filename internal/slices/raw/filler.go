package raw

type Filler struct {
	data Tags

	keys map[string]struct{}
}

func NewFiller() *Filler {
	return &Filler{keys: map[string]struct{}{}}
}

func (f *Filler) Data() Tags {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	if _, ok := f.keys[key]; ok {
		// Ignore duplicated keys.
		// TODO(ldez) add an option to through an error.
		return nil
	}

	f.keys[key] = struct{}{}

	f.data = append(f.data, Tag{
		Key:   key,
		Value: value,
	})

	return nil
}
