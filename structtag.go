package structtags

import (
	"github.com/fatih/structtag"
	mapsmultikeys "github.com/ldez/structtags/internal/maps/multikeys"
	mapsraw "github.com/ldez/structtags/internal/maps/raw"
	mapsvalues "github.com/ldez/structtags/internal/maps/values"
	slicesfatih "github.com/ldez/structtags/internal/slices/fatih"
	sliceraw "github.com/ldez/structtags/internal/slices/raw"
	slicevalues "github.com/ldez/structtags/internal/slices/values"
	"github.com/ldez/structtags/structured"
)

// ParseToMap parses a struct tag to a `map[string]string`.
// Ignore duplicated keys by default.
func ParseToMap(tag string, options ...mapsraw.Option) (mapsraw.Tag, error) {
	return mapsraw.Parse(tag, options...)
}

// ParseToMapMultikeys parses a struct tag to a `map[string][]string`.
// For non-conventional tags where the key is repeated.
func ParseToMapMultikeys(tag string) (mapsmultikeys.Tag, error) {
	return mapsmultikeys.Parse(tag)
}

// ParseToMapValues parses a struct tag to a `map[string][]string`.
// The value is split on comma.
// Ignore duplicated keys by default.
func ParseToMapValues(tag string, options ...mapsvalues.Option) (mapsvalues.Tag, error) {
	return mapsvalues.Parse(tag, options...)
}

// ParseToSlice parses a struct tag to a slice of [sliceraw.Tag].
// Ignore duplicated keys by default.
func ParseToSlice(tag string, options ...sliceraw.Option) (sliceraw.Tags, error) {
	return sliceraw.Parse(tag, options...)
}

// ParseToSliceValues parses a struct tag to a slice of [slicevalues.Tag].
// The value is split on comma.
// Ignore duplicated keys by default.
func ParseToSliceValues(tag string, options ...slicevalues.Option) (slicevalues.Tags, error) {
	return slicevalues.Parse(tag, options...)
}

// ParseToSliceStructured parses a struct tag to a [structured.Tag].
// Allows modifying the struct tags.
// The value is split on comma.
// Ignore duplicated keys by default.
func ParseToSliceStructured(tag string, options ...structured.Option) (*structured.Tag, error) {
	return structured.Parse(tag, options...)
}

// ParseToFatih parses a struct tag to a [*structtag.Tags].
// The value is split on comma.
func ParseToFatih(tag string, escapeComma bool) (*structtag.Tags, error) {
	return slicesfatih.Parse(tag, escapeComma)
}
