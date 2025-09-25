package parser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestTag struct {
	Key   string
	Value string
}

type TestFiller struct {
	data []TestTag
}

func (f *TestFiller) Data() []TestTag {
	return f.data
}

func (f *TestFiller) Fill(key, value string) error {
	if key == "oops" {
		return errors.New("oops")
	}

	f.data = append(f.data, TestTag{
		Key:   key,
		Value: value,
	})

	return nil
}

func TestParseTag(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      string
		expected []TestTag
	}{
		{
			desc:     "no tag",
			tag:      "",
			expected: nil,
		},
		{
			desc:     "empty value",
			tag:      `json:""`,
			expected: []TestTag{{Key: "json", Value: ""}},
		},
		{
			desc:     "whitespace value",
			tag:      ` `,
			expected: nil,
		},
		{
			desc:     "simple value",
			tag:      `json:"a"`,
			expected: []TestTag{{Key: "json", Value: "a"}},
		},
		{
			desc:     "simple value (double quotes)",
			tag:      "json:\"a\"",
			expected: []TestTag{{Key: "json", Value: "a"}},
		},
		{
			desc:     "multiple values",
			tag:      `json:"a,b,c"`,
			expected: []TestTag{{Key: "json", Value: "a,b,c"}},
		},
		{
			desc:     "quoted value",
			tag:      `json:"a:\"b\""`,
			expected: []TestTag{{Key: "json", Value: "a:\"b\""}},
		},
		{
			desc: "multiple empty tag",
			tag:  `json:"" yaml:""`,
			expected: []TestTag{
				{Key: "json", Value: ""},
				{Key: "yaml", Value: ""},
			},
		},
		{
			desc: "multiple tag",
			tag:  `json:"a" yaml:"b"`,
			expected: []TestTag{
				{Key: "json", Value: "a"},
				{Key: "yaml", Value: "b"},
			},
		},
		{
			desc: "identical tag",
			tag:  `json:"a" json:"b"`,
			expected: []TestTag{
				{Key: "json", Value: "a"},
				{Key: "json", Value: "b"},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			tags, err := Tag(test.tag, &TestFiller{})
			require.NoError(t, err)

			assert.Equal(t, test.expected, tags)
		})
	}
}

func TestParseTag_error(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      string
		expected string
	}{
		{
			desc:     "invalid tag",
			tag:      `json`,
			expected: "syntax error in tag \"json\"",
		},
		{
			desc:     "incomplete tag",
			tag:      `json:"`,
			expected: "syntax error in tag \"\\\"\"",
		},
		{
			desc:     "filler error",
			tag:      `oops:""`,
			expected: "oops",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			_, err := Tag(test.tag, &TestFiller{})
			require.Error(t, err)

			assert.EqualError(t, err, test.expected)
		})
	}
}
