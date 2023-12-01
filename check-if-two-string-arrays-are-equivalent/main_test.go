package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestArrayStringsAreEqual(t *testing.T) {
	tt := []struct {
		Word1    []string
		Word2    []string
		Expected bool
	}{
		{
			Word1:    []string{"ab", "c"},
			Word2:    []string{"a", "bc"},
			Expected: true,
		},
		{
			Word1:    []string{"a", "cb"},
			Word2:    []string{"ab", "c"},
			Expected: false,
		},
		{
			Word1:    []string{"abc", "d", "defg"},
			Word2:    []string{"abcddefg"},
			Expected: true,
		},
		{
			Word1:    []string{"abc", "d", "defg"},
			Word2:    []string{"abcddef"},
			Expected: false,
		},
		{
			Word1:    []string{"ab", "c"},
			Word2:    []string{"a", "bcd"},
			Expected: false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, arrayStringsAreEqual(tc.Word1, tc.Word2))
	}
}
