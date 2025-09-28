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
// Ignore duplicated keys.
func ParseToMap(tag string) (mapsraw.Tag, error) {
	return mapsraw.Parse(tag)
}

// ParseToMapMultikeys parses a struct tag to a `map[string][]string`.
// For non-conventional tags where the key is repeated.
func ParseToMapMultikeys(tag string) (mapsmultikeys.Tag, error) {
	return mapsmultikeys.Parse(tag)
}

// ParseToMapValues parses a struct tag to a `map[string][]string`.
// The value is split on comma.
// Ignore duplicated keys.
func ParseToMapValues(tag string, escapeComma bool) (mapsvalues.Tag, error) {
	return mapsvalues.Parse(tag, escapeComma)
}

// ParseToSlice parses a struct tag to a slice of [sliceraw.Tag].
func ParseToSlice(tag string) (sliceraw.Tags, error) {
	return sliceraw.Parse(tag)
}

// ParseToSliceValues parses a struct tag to a slice of [slicevalues.Tag].
// The value is split on comma.
func ParseToSliceValues(tag string, escapeComma bool) (slicevalues.Tags, error) {
	return slicevalues.Parse(tag, escapeComma)
}

// ParseToSliceStructured parses a struct tag to a [structured.Tag].
// Allows modifying the struct tags.
// The value is split on comma.
func ParseToSliceStructured(tag string, opt *structured.Options) (*structured.Tag, error) {
	return structured.Parse(tag, opt)
}

// ParseToFatih parses a struct tag to a [*structtag.Tags].
// The value is split on comma.
func ParseToFatih(tag string, escapeComma bool) (*structtag.Tags, error) {
	return slicesfatih.Parse(tag, escapeComma)
}
