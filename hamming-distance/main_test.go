package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestHammingDistance(t *testing.T) {
	tt := []struct {
		X        int
		Y        int
		Expected int
	}{
		{
			X:        1,
			Y:        4,
			Expected: 2,
		},
		{
			X:        3,
			Y:        1,
			Expected: 1,
		},
		{
			X:        58,
			Y:        405,
			Expected: 7,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, hammingDistance(tc.X, tc.Y), tc.Expected)
	}
}
