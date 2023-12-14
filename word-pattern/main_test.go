package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordPattern(t *testing.T) {
	tt := []struct {
		Pattern  string
		S        string
		Expected bool
	}{
		{
			Pattern:  "abba",
			S:        "dog cat cat dog",
			Expected: true,
		},
		{
			Pattern:  "abba",
			S:        "dog cat cat fish",
			Expected: false,
		},
		{
			Pattern:  "aaaa",
			S:        "dog cat cat dog",
			Expected: false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, wordPattern(tc.Pattern, tc.S))
	}
}
