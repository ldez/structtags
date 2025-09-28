package raw

import (
	"fmt"
	"strings"
)

type Tags []Tag

func (t Tags) String() string {
	var b strings.Builder

	for _, e := range t {
		b.WriteString(fmt.Sprintf("%s:%q ", e.Key, e.Value))
	}

	return strings.TrimSuffix(b.String(), " ")
}

type Tag struct {
	Key   string
	Value string
}
