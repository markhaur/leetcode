package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLongestPalindrome(t *testing.T) {
	tt := []struct {
		S        string
		Expected int
	}{
		{
			S:        "abccccdd",
			Expected: 7,
		},
		{
			S:        "a",
			Expected: 1,
		},
		{
			S:        "abccccddddd",
			Expected: 9,
		},
		{
			S:        "abcccckfdihfnxkcnhdfdfbjdndmfnvjbnASDFDFDFDFDDFddddd",
			Expected: 41,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, longestPalindrome(tc.S), tc.Expected)
	}
}
