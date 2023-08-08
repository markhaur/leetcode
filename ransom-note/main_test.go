package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	tt := []struct {
		Magazine   string
		RansomNote string
		Expected   bool
	}{
		{
			Magazine:   "a",
			RansomNote: "b",
			Expected:   false,
		},
		{
			Magazine:   "ab",
			RansomNote: "aa",
			Expected:   false,
		},
		{
			Magazine:   "aab",
			RansomNote: "aa",
			Expected:   true,
		},
		{
			Magazine:   "baa",
			RansomNote: "aab",
			Expected:   true,
		},
		{
			Magazine:   "baca",
			RansomNote: "aab",
			Expected:   true,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, canConstruct(tc.RansomNote, tc.Magazine), tc.Expected)
	}
}
