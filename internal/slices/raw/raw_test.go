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

	expected := Tags{
		{Key: "a", Value: "b"},
		{Key: "d", Value: "e,f\\,g"},
	}

	assert.Equal(t, expected, filler.Data())
}

func TestFiller_Fill_duplicate(t *testing.T) {
	filler := Filler{}

	err := filler.Fill("a", "b")
	require.NoError(t, err)

	err = filler.Fill("a", "c")
	require.NoError(t, err)

	expected := Tags{
		{Key: "a", Value: "b"},
		{Key: "a", Value: "c"},
	}

	assert.Equal(t, expected, filler.Data())
}

func TestTas_String(t *testing.T) {
	testCases := []struct {
		desc     string
		tag      Tags
		expected string
	}{
		{
			desc:     "empty",
			expected: "",
		},
		{
			desc:     "one entry",
			tag:      Tags{{Key: "a", Value: "b"}},
			expected: `a:"b"`,
		},
		{
			desc: "multiple entries",
			tag: Tags{
				{Key: "a", Value: "b"},
				{Key: "c", Value: "d,e"},
			},
			expected: `a:"b" c:"d,e"`,
		},
		{
			desc: "duplicate entry",
			tag: Tags{
				{Key: "a", Value: "b"},
				{Key: "a", Value: "d,e"},
			},
			expected: `a:"b" a:"d,e"`,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.expected, test.tag.String())
		})
	}
}
