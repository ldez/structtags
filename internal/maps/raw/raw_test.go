package raw

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFiller_Fill(t *testing.T) {
	filler := Filler{}

	err := filler.Fill("a", "b")
	require.NoError(t, err)

	err = filler.Fill("d", "e,f\\,g")
	require.NoError(t, err)

	expected := Tag{
		"a": "b",
		"d": "e,f\\,g",
	}

	assert.Equal(t, expected, filler.Data())
}

func TestFiller_Fill_duplicate(t *testing.T) {
	filler := Filler{}

	err := filler.Fill("a", "b")
	require.NoError(t, err)

	err = filler.Fill("a", "c")
	require.NoError(t, err)

	expected := Tag{"a": "b"}

	assert.Equal(t, expected, filler.Data())
}

func TestTag_String(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      Tag
		expected string
	}{
		{
			desc:     "empty",
			expected: "",
		},
		{
			desc:     "one entry",
			tag:      Tag{"a": "b"},
			expected: `a:"b"`,
		},
		{
			desc: "multiple entries",
			tag: Tag{
				"a": "b",
				"c": "d,e",
			},
			expected: `a:"b" c:"d,e"`,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.expected, test.tag.String())
		})
	}
}
