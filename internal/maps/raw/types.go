package raw

import (
	"fmt"
	"strings"
)

// Tag is a key/value map.
type Tag map[string]string

func (m Tag) String() string {
	var b strings.Builder

	for k, v := range m {
		b.WriteString(fmt.Sprintf("%s:%q ", k, v))
	}

	return strings.TrimSuffix(b.String(), " ")
}
