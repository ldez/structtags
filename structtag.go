package structtags

import (
	"github.com/fatih/structtag"
	mapsmultikeys "github.com/ldez/structtags/internal/maps/multikeys"
	mapsraw "github.com/ldez/structtags/internal/maps/raw"
	mapsvalues "github.com/ldez/structtags/internal/maps/values"
	slicesfatih "github.com/ldez/structtags/internal/slices/fatih"
	sliceraw "github.com/ldez/structtags/internal/slices/raw"
	slicevalues "github.com/ldez/structtags/internal/slices/values"
	"github.com/ldez/structtags/parser"
	"github.com/ldez/structtags/structured"
)

// ParseToMap parses a struct tag to a `map[string]string`.
func ParseToMap(tag string) (mapsraw.Tag, error) {
	return parser.Tag(tag, &mapsraw.Filler{})
}

// ParseToMapMultikeys parses a struct tag to a `map[string][]string`.
// For non-conventional tags where the key is repeated.
func ParseToMapMultikeys(tag string) (mapsmultikeys.Tag, error) {
	return parser.Tag(tag, &mapsmultikeys.Filler{})
}

// ParseToMapValues parses a struct tag to a `map[string][]string`.
// The value is split on comma.
func ParseToMapValues(tag string, escapeComma bool) (mapsvalues.Tag, error) {
	return parser.Tag(tag, mapsvalues.NewFiller(escapeComma))
}

// ParseToSlice parses a struct tag to a slice of [sliceraw.Tag].
func ParseToSlice(tag string) (sliceraw.Tags, error) {
	return parser.Tag(tag, &sliceraw.Filler{})
}

// ParseToSliceValues parses a struct tag to a slice of [slicevalues.Tag].
// The value is split on comma.
func ParseToSliceValues(tag string, escapeComma bool) (slicevalues.Tags, error) {
	return parser.Tag(tag, slicevalues.NewFiller(escapeComma))
}

// ParseToSliceStructured parses a struct tag to a [structured.Tag].
// Allows modifying the struct tags.
// The value is split on comma.
func ParseToSliceStructured(tag string, opt *structured.Options) (*structured.Tag, error) {
	var escapeComma bool

	var allowDuplicateKeys bool

	if opt != nil {
		escapeComma = opt.EscapeComma
		allowDuplicateKeys = opt.AllowDuplicateKeys
	}

	if tag == "" {
		return structured.NewTag(escapeComma, allowDuplicateKeys), nil
	}

	return parser.Tag(tag, structured.NewFiller(escapeComma, allowDuplicateKeys))
}

// ParseToFatih parses a struct tag to a [*structtag.Tags].
// The value is split on comma.
func ParseToFatih(tag string, escapeComma bool) (*structtag.Tags, error) {
	tags, err := parser.Tag(tag, slicesfatih.NewFiller(escapeComma))
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
