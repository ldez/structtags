package multikeys

// Tag is a key/values map.
// There is no `String` method because
// it's not possible to differentiate values from a key and values from a repeated key.
type Tag map[string][]string
