package structured

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type tuple struct {
	key, value string
}

func TestFiller_Fill(t *testing.T) {
	testCases := []struct {
		desc               string
		data               []tuple
		escapeComma        bool
		allowDuplicateKeys bool
		expected           *Tag
	}{
		{
			desc: "option: no escape, no duplicate",
			data: []tuple{
				{key: "a", value: "b"},
				{key: "d", value: "e"},
			},
			escapeComma:        false,
			allowDuplicateKeys: false,
			expected: &Tag{
				entries: []*Entry{
					{
						Key:         "a",
						RawValue:    "b",
						escapeComma: false,
					},
					{
						Key:         "d",
						RawValue:    "e",
						escapeComma: false,
					},
				},
				escapeComma:        false,
				allowDuplicateKeys: false,
			},
		},
		{
			desc: "option: escape, no duplicate",
			data: []tuple{
				{key: "a", value: "b"},
				{key: "d", value: "e"},
			},
			escapeComma:        true,
			allowDuplicateKeys: false,
			expected: &Tag{
				entries: []*Entry{
					{
						Key:         "a",
						RawValue:    "b",
						escapeComma: true,
					},
					{
						Key:         "d",
						RawValue:    "e",
						escapeComma: true,
					},
				},
				escapeComma:        true,
				allowDuplicateKeys: false,
			},
		},
		{
			desc: "option: no escape, duplicate",
			data: []tuple{
				{key: "a", value: "b"},
				{key: "d", value: "e"},
			},
			escapeComma:        false,
			allowDuplicateKeys: true,
			expected: &Tag{
				entries: []*Entry{
					{
						Key:         "a",
						RawValue:    "b",
						escapeComma: false,
					},
					{
						Key:         "d",
						RawValue:    "e",
						escapeComma: false,
					},
				},
				escapeComma:        false,
				allowDuplicateKeys: true,
			},
		},
		{
			desc: "option: escape, duplicate",
			data: []tuple{
				{key: "a", value: "b"},
				{key: "d", value: "e"},
			},
			escapeComma:        true,
			allowDuplicateKeys: true,
			expected: &Tag{
				entries: []*Entry{
					{
						Key:         "a",
						RawValue:    "b",
						escapeComma: true,
					},
					{
						Key:         "d",
						RawValue:    "e",
						escapeComma: true,
					},
				},
				escapeComma:        true,
				allowDuplicateKeys: true,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			filler := NewFiller(test.escapeComma, test.allowDuplicateKeys)

			for _, d := range test.data {
				require.NoError(t, filler.Fill(d.key, d.value))
			}

			assert.Equal(t, test.expected, filler.Data())
		})
	}
}

func TestFiller_Fill_no_duplicate(t *testing.T) {
	filler := NewFiller(false, false)

	err := filler.Fill("a", "b")
	require.NoError(t, err)

	err = filler.Fill("a", "c")
	require.Error(t, err)
}

func TestFiller_Fill_duplicate(t *testing.T) {
	filler := NewFiller(false, true)

	err := filler.Fill("a", "b")
	require.NoError(t, err)

	err = filler.Fill("a", "c")
	require.NoError(t, err)

	expected := &Tag{
		entries: []*Entry{
			{Key: "a", RawValue: "b"},
			{Key: "a", RawValue: "c"},
		},
		escapeComma:        false,
		allowDuplicateKeys: true,
	}

	assert.Equal(t, expected, filler.Data())
}
