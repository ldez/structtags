package raw

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFiller_Fill(t *testing.T) {
	filler := NewFiller()

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
	filler := NewFiller()

	err := filler.Fill("a", "b")
	require.NoError(t, err)

	err = filler.Fill("a", "c")
	require.NoError(t, err)

	expected := Tags{
		{Key: "a", Value: "b"},
	}

	assert.Equal(t, expected, filler.Data())
}
