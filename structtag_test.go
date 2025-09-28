package structtags

import (
	"slices"
	"testing"

	"github.com/fatih/structtag"
	sliceraw "github.com/ldez/structtags/internal/slices/raw"
	slicevalues "github.com/ldez/structtags/internal/slices/values"
	"github.com/ldez/structtags/structured"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseToMap(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      string
		expected map[string]string
	}{
		{
			desc:     "no tag",
			tag:      "",
			expected: nil,
		},
		{
			desc:     "empty value",
			tag:      `json:""`,
			expected: map[string]string{"json": ""},
		},
		{
			desc:     "simple value",
			tag:      `json:"a"`,
			expected: map[string]string{"json": "a"},
		},
		{
			desc:     "multiple values",
			tag:      `json:"a,b,c"`,
			expected: map[string]string{"json": "a,b,c"},
		},
		{
			desc:     "quoted value",
			tag:      `json:"a:\"b\""`,
			expected: map[string]string{"json": "a:\"b\""},
		},
		{
			desc: "multiple empty tag",
			tag:  `json:"" yaml:""`,
			expected: map[string]string{
				"json": "",
				"yaml": "",
			},
		},
		{
			desc: "multiple tag",
			tag:  `json:"a" yaml:"b"`,
			expected: map[string]string{
				"json": "a",
				"yaml": "b",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			tags, err := ParseToMap(test.tag)
			require.NoError(t, err)

			assert.Equal(t, test.expected, tags)
		})
	}
}

func TestParseToMapMultikeys(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      string
		expected map[string][]string
	}{
		{
			desc:     "no tag",
			tag:      "",
			expected: nil,
		},
		{
			desc:     "empty value",
			tag:      `json:""`,
			expected: map[string][]string{"json": {""}},
		},
		{
			desc:     "simple value",
			tag:      `json:"a"`,
			expected: map[string][]string{"json": {"a"}},
		},
		{
			desc:     "multiple values",
			tag:      `json:"a,b,c"`,
			expected: map[string][]string{"json": {"a,b,c"}},
		},
		{
			desc:     "quoted value",
			tag:      `json:"a:\"b\""`,
			expected: map[string][]string{"json": {"a:\"b\""}},
		},
		{
			desc:     "multiple empty tag",
			tag:      `json:"" yaml:""`,
			expected: map[string][]string{"json": {""}, "yaml": {""}},
		},
		{
			desc:     "multiple tag",
			tag:      `json:"a" yaml:"b"`,
			expected: map[string][]string{"json": {"a"}, "yaml": {"b"}},
		},
		{
			desc:     "identical keys",
			tag:      `json:"a" json:"b"`,
			expected: map[string][]string{"json": {"a", "b"}},
		},
		{
			desc: "foo",
			tag:  `long:"thresholds" default:"1" default:"2" env:"THRESHOLD_VALUES"  env-delim:","`,
			expected: map[string][]string{
				"default":   {"1", "2"},
				"env":       {"THRESHOLD_VALUES"},
				"env-delim": {","},
				"long":      {"thresholds"},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			tags, err := ParseToMapMultikeys(test.tag)
			require.NoError(t, err)

			assert.Equal(t, test.expected, tags)
		})
	}
}

func TestParseToMapValues(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      string
		expected map[string][]string
	}{
		{
			desc:     "no tag",
			tag:      "",
			expected: nil,
		},
		{
			desc:     "empty value",
			tag:      `json:""`,
			expected: map[string][]string{"json": {""}},
		},
		{
			desc:     "simple value",
			tag:      `json:"a"`,
			expected: map[string][]string{"json": {"a"}},
		},
		{
			desc:     "multiple values",
			tag:      `json:"a,b,c"`,
			expected: map[string][]string{"json": {"a", "b", "c"}},
		},
		{
			desc:     "quoted value",
			tag:      `json:"a:\"b\""`,
			expected: map[string][]string{"json": {"a:\"b\""}},
		},
		{
			desc:     "escaped coma",
			tag:      `json:"b\\,c\\,d,e"`,
			expected: map[string][]string{"json": {"b\\,c\\,d", "e"}},
		},
		{
			desc:     "multiple empty tag",
			tag:      `json:"" yaml:""`,
			expected: map[string][]string{"json": {""}, "yaml": {""}},
		},
		{
			desc:     "multiple tag",
			tag:      `json:"a" yaml:"b"`,
			expected: map[string][]string{"json": {"a"}, "yaml": {"b"}},
		},
		{
			desc:     "identical keys",
			tag:      `json:"a" json:"b"`,
			expected: map[string][]string{"json": {"a", "b"}},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			tags, err := ParseToMapValues(test.tag, true)
			require.NoError(t, err)

			assert.Equal(t, test.expected, tags)
		})
	}
}

func TestParseToSlice(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      string
		expected []sliceraw.Tag
	}{
		{
			desc:     "no tag",
			tag:      "",
			expected: nil,
		},
		{
			desc:     "empty value",
			tag:      `json:""`,
			expected: []sliceraw.Tag{{Key: "json", Value: ""}},
		},
		{
			desc:     "simple value",
			tag:      `json:"a"`,
			expected: []sliceraw.Tag{{Key: "json", Value: "a"}},
		},
		{
			desc:     "multiple values",
			tag:      `json:"a,b,c"`,
			expected: []sliceraw.Tag{{Key: "json", Value: "a,b,c"}},
		},
		{
			desc:     "quoted value",
			tag:      `json:"a:\"b\""`,
			expected: []sliceraw.Tag{{Key: "json", Value: "a:\"b\""}},
		},
		{
			desc: "multiple empty tag",
			tag:  `json:"" yaml:""`,
			expected: []sliceraw.Tag{
				{Key: "json", Value: ""},
				{Key: "yaml", Value: ""},
			},
		},
		{
			desc: "multiple tag",
			tag:  `json:"a" yaml:"b"`,
			expected: []sliceraw.Tag{
				{Key: "json", Value: "a"},
				{Key: "yaml", Value: "b"},
			},
		},
		{
			desc: "identical keys",
			tag:  `json:"a" json:"b"`,
			expected: []sliceraw.Tag{
				{Key: "json", Value: "a"},
				{Key: "json", Value: "b"},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			tags, err := ParseToSlice(test.tag)
			require.NoError(t, err)

			assert.Equal(t, test.expected, tags)
		})
	}
}

func TestParseToSliceValues(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      string
		expected []slicevalues.Tag
	}{
		{
			desc:     "no tag",
			tag:      "",
			expected: nil,
		},
		{
			desc:     "empty value",
			tag:      `json:""`,
			expected: []slicevalues.Tag{{Key: "json", Values: []string{""}}},
		},
		{
			desc:     "simple value",
			tag:      `json:"a"`,
			expected: []slicevalues.Tag{{Key: "json", Values: []string{"a"}}},
		},
		{
			desc:     "multiple values",
			tag:      `json:"a,b,c"`,
			expected: []slicevalues.Tag{{Key: "json", Values: []string{"a", "b", "c"}}},
		},
		{
			desc:     "quoted value",
			tag:      `json:"a:\"b\""`,
			expected: []slicevalues.Tag{{Key: "json", Values: []string{"a:\"b\""}}},
		},
		{
			desc:     "escaped coma",
			tag:      `json:"b\\,c\\,d,e"`,
			expected: []slicevalues.Tag{{Key: "json", Values: []string{"b\\,c\\,d", "e"}}},
		},
		{
			desc: "multiple empty tag",
			tag:  `json:"" yaml:""`,
			expected: []slicevalues.Tag{
				{Key: "json", Values: []string{""}},
				{Key: "yaml", Values: []string{""}},
			},
		},
		{
			desc: "multiple tag",
			tag:  `json:"a" yaml:"b"`,
			expected: []slicevalues.Tag{
				{Key: "json", Values: []string{"a"}},
				{Key: "yaml", Values: []string{"b"}},
			},
		},
		{
			desc: "identical keys",
			tag:  `json:"a" json:"b"`,
			expected: []slicevalues.Tag{
				{Key: "json", Values: []string{"a"}},
				{Key: "json", Values: []string{"b"}},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			tags, err := ParseToSliceValues(test.tag, true)
			require.NoError(t, err)

			assert.Equal(t, test.expected, tags)
		})
	}
}

func TestParseToFatih(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      string
		expected []*structtag.Tag
	}{
		{
			desc:     "no tag",
			tag:      "",
			expected: nil,
		},
		{
			desc: "empty value",
			tag:  `json:""`,
			expected: []*structtag.Tag{
				{Key: "json", Name: ""},
			},
		},
		{
			desc: "simple value",
			tag:  `json:"a"`,
			expected: []*structtag.Tag{
				{Key: "json", Name: "a"},
			},
		},
		{
			desc: "multiple values",
			tag:  `json:"a,b,c"`,
			expected: []*structtag.Tag{{
				Key:     "json",
				Name:    "a",
				Options: []string{"b", "c"},
			}},
		},
		{
			desc: "quoted value",
			tag:  `json:"a:\"b\""`,
			expected: []*structtag.Tag{
				{Key: "json", Name: "a:\"b\""},
			},
		},
		{
			desc: "escaped coma",
			tag:  `json:"b\\,c\\,d,e"`,
			expected: []*structtag.Tag{{
				Key:     "json",
				Name:    "b\\",
				Options: []string{"c\\", "d", "e"},
			}},
		},
		{
			desc: "multiple empty tag",
			tag:  `json:"" yaml:""`,
			expected: []*structtag.Tag{
				{Key: "json", Name: ""},
				{Key: "yaml", Name: ""},
			},
		},
		{
			desc: "multiple tag",
			tag:  `json:"a" yaml:"b"`,
			expected: []*structtag.Tag{
				{Key: "json", Name: "a"},
				{Key: "yaml", Name: "b"},
			},
		},
		{
			desc: "identical keys",
			tag:  `json:"a" json:"b"`,
			expected: []*structtag.Tag{
				{Key: "json", Name: "b"},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			tags, err := ParseToFatih(test.tag, false)
			require.NoError(t, err)

			if test.expected == nil {
				if !assert.Nil(t, tags) {
					assert.Equal(t, test.expected, tags.Tags())
				}
			} else {
				assert.Equal(t, test.expected, tags.Tags())
			}
		})
	}
}

func TestParseToFatih_extended(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      string
		expected []*structtag.Tag
	}{
		{
			desc:     "no tag",
			tag:      "",
			expected: nil,
		},
		{
			desc: "empty value",
			tag:  `json:""`,
			expected: []*structtag.Tag{
				{Key: "json", Name: ""},
			},
		},
		{
			desc: "simple value",
			tag:  `json:"a"`,
			expected: []*structtag.Tag{
				{Key: "json", Name: "a"},
			},
		},
		{
			desc: "multiple values",
			tag:  `json:"a,b,c"`,
			expected: []*structtag.Tag{{
				Key:     "json",
				Name:    "a",
				Options: []string{"b", "c"},
			}},
		},
		{
			desc: "quoted value",
			tag:  `json:"a:\"b\""`,
			expected: []*structtag.Tag{
				{Key: "json", Name: "a:\"b\""},
			},
		},
		{
			desc: "escaped coma",
			tag:  `json:"b\\,c\\,d,e"`,
			expected: []*structtag.Tag{{
				Key:     "json",
				Name:    "b\\,c\\,d",
				Options: []string{"e"},
			}},
		},
		{
			desc: "multiple empty tag",
			tag:  `json:"" yaml:""`,
			expected: []*structtag.Tag{
				{Key: "json", Name: ""},
				{Key: "yaml", Name: ""},
			},
		},
		{
			desc: "multiple tag",
			tag:  `json:"a" yaml:"b"`,
			expected: []*structtag.Tag{
				{Key: "json", Name: "a"},
				{Key: "yaml", Name: "b"},
			},
		},
		{
			desc: "identical keys",
			tag:  `json:"a" json:"b"`,
			expected: []*structtag.Tag{
				{Key: "json", Name: "b"},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			tags, err := ParseToFatih(test.tag, true)
			require.NoError(t, err)

			if test.expected == nil {
				if !assert.Nil(t, tags) {
					assert.Equal(t, test.expected, tags.Tags())
				}
			} else {
				assert.Equal(t, test.expected, tags.Tags())
			}
		})
	}
}

func TestParseToSliceStructured(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      string
		options  *structured.Options
		expected []*structured.Entry
	}{
		{
			desc:     "no tag",
			tag:      "",
			expected: nil,
		},
		{
			desc: "empty value",
			tag:  `json:""`,
			expected: []*structured.Entry{
				{Key: "json", RawValue: ""},
			},
		},
		{
			desc: "simple value",
			tag:  `json:"a"`,
			expected: []*structured.Entry{
				{Key: "json", RawValue: "a"},
			},
		},
		{
			desc: "multiple values",
			tag:  `json:"a,b,c"`,

			expected: []*structured.Entry{
				{Key: "json", RawValue: "a,b,c"},
			},
		},
		{
			desc: "quoted value",
			tag:  `json:"a:\"b\""`,
			expected: []*structured.Entry{
				{Key: "json", RawValue: "a:\"b\""},
			},
		},
		{
			desc: "ignore escaped coma",
			tag:  `json:"b\\,c\\,d,e"`,
			expected: []*structured.Entry{
				{Key: "json", RawValue: "b\\,c\\,d,e"},
			},
		},
		{
			desc: "multiple empty tag",
			tag:  `json:"" yaml:""`,
			expected: []*structured.Entry{
				{Key: "json", RawValue: ""},
				{Key: "yaml", RawValue: ""},
			},
		},
		{
			desc: "multiple tag",
			tag:  `json:"a" yaml:"b"`,
			expected: []*structured.Entry{
				{Key: "json", RawValue: "a"},
				{Key: "yaml", RawValue: "b"},
			},
		},
		{
			desc:    "identical keys",
			tag:     `json:"a" json:"b"`,
			options: &structured.Options{AllowDuplicateKeys: true},
			expected: []*structured.Entry{
				{Key: "json", RawValue: "a"},
				{Key: "json", RawValue: "b"},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			tags, err := ParseToSliceStructured(test.tag, test.options)
			require.NoError(t, err)

			if test.expected == nil {
				if !assert.True(t, tags.IsEmpty()) {
					assert.Equal(t, test.expected, slices.Collect(tags.Seq()))
				}
			} else {
				assert.Equal(t, test.expected, slices.Collect(tags.Seq()))
			}
		})
	}
}
