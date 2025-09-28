package raw

type Filler struct {
	data map[string]string
}

func (f *Filler) Data() map[string]string {
	return f.data
}

func (f *Filler) Fill(key, value string) error {
	if f.data != nil && f.data[key] != "" {
		// Ignore duplicated key.
		// TODO(ldez) add an option to through an error.
		return nil
	}

	if f.data == nil {
		f.data = map[string]string{}
	}

	f.data[key] = value

	return nil
}
