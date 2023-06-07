package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		value    string
		expected bool
	}{
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
		{
			" ",
			true,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, isPalindrome(tc.value), tc.expected)
	}
}
