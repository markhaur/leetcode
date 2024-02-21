package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeBitwiseAnd(t *testing.T) {
	tt := []struct {
		Left     int
		Right    int
		Expected int
	}{
		{
			Left:     5,
			Right:    7,
			Expected: 4,
		},
		{
			Left:     0,
			Right:    0,
			Expected: 0,
		},
		{
			Left:     1,
			Right:    2147483647,
			Expected: 0,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, rangeBitwiseAnd(tc.Left, tc.Right))
	}
}
