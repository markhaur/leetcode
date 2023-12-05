package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfMatches(t *testing.T) {
	tt := []struct {
		N        int
		Expected int
	}{
		{
			N:        7,
			Expected: 6,
		},
		{
			N:        14,
			Expected: 13,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, numberOfMatches(tc.N))
	}
}