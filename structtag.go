package structtags

import (
	"github.com/fatih/structtag"
	mapsmultikeys "github.com/ldez/structtags/internal/maps/multikeys"
	mapsraw "github.com/ldez/structtags/internal/maps/raw"
	mapsvalues "github.com/ldez/structtags/internal/maps/values"
	slicesfatih "github.com/ldez/structtags/internal/slices/fatih"
	slicesfatihext "github.com/ldez/structtags/internal/slices/fatihext"
	sliceraw "github.com/ldez/structtags/internal/slices/raw"
	slicevalues "github.com/ldez/structtags/internal/slices/values"
	"github.com/ldez/structtags/parser"
)

// ParseToMap parses a struct tag to a `map[string]string`.
func ParseToMap(tag string) (map[string]string, error) {
	return parser.Tag(tag, &mapsraw.Filler{})
}

// ParseToMapMultikeys parses a struct tag to a `map[string][]string`.
// For non-conventional tags where the key is repeated.
func ParseToMapMultikeys(tag string) (map[string][]string, error) {
	return parser.Tag(tag, &mapsmultikeys.Filler{})
}

// ParseToMapValues parses a struct tag to a `map[string][]string`.
// The value is split on comma.
func ParseToMapValues(tag string) (map[string][]string, error) {
	return parser.Tag(tag, &mapsvalues.Filler{})
}

// ParseToSlice parses a struct tag to a slice of [parser.Tag].
func ParseToSlice(tag string) ([]sliceraw.Tag, error) {
	return parser.Tag(tag, &sliceraw.Filler{})
}

// ParseToSliceValues parses a struct tag to a slice of [slicevalues.Tag].
// The value is split on comma.
func ParseToSliceValues(tag string) ([]slicevalues.Tag, error) {
	return parser.Tag(tag, &slicevalues.Filler{})
}

// ParseToFatih parses a struct tag to a [*structtag.Tags].
// The value is split on comma.
func ParseToFatih(tag string) (*structtag.Tags, error) {
	tags, err := parser.Tag(tag, &slicesfatih.Filler{})
	if err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return nil, nil
	}

	ftgs := &structtag.Tags{}

	for _, s := range tags {
		if err := ftgs.Set(s); err != nil {
			return nil, err
		}
	}

	return ftgs, err
}

// ParseToFatihExtended parses a struct tag to a [*structtag.Tags].
// Extended: support comma escaped by backslash.
func ParseToFatihExtended(tag string) (*structtag.Tags, error) {
	tags, err := parser.Tag(tag, &slicesfatihext.Filler{})
	if err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return nil, nil
	}

	ftgs := &structtag.Tags{}

	for _, s := range tags {
		if err := ftgs.Set(s); err != nil {
			return nil, err
		}
	}

	return ftgs, err
}
