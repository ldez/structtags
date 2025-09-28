package values

import (
	"fmt"
	"strings"
)

// Tag is a key/values map.
type Tag map[string][]string

func (m Tag) String() string {
	var b strings.Builder

	for k, v := range m {
		b.WriteString(fmt.Sprintf("%s:%q ", k, strings.Join(v, ",")))
	}

	return strings.TrimSuffix(b.String(), " ")
}
