package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWinNim(t *testing.T) {
	tt := []struct {
		N        int
		Expected bool
	}{
		{
			N:        4,
			Expected: false,
		},
		{
			N:        1,
			Expected: true,
		},
		{
			N:        2,
			Expected: true,
		},
		{
			N:        5,
			Expected: true,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, canWinNim(tc.N))
	}
}
