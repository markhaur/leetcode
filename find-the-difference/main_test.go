package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheDifference(t *testing.T) {
	tt := []struct {
		S        string
		T        string
		Expected byte
	}{
		{
			S:        "abcd",
			T:        "abcde",
			Expected: 'e',
		},
		{
			S:        "",
			T:        "y",
			Expected: 'y',
		},
		{
			S:        "abcd",
			T:        "adcbe",
			Expected: 'e',
		},
	}

	for _, tc := range tt {
		assert.Equal(t, findTheDifference(tc.S, tc.T), tc.Expected)
	}
}
