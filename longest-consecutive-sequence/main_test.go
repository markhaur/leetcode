package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	tt := []struct {
		Nums     []int
		Expected int
	}{
		{
			Nums:     []int{100, 4, 200, 1, 3, 2},
			Expected: 4,
		},
		{
			Nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			Expected: 9,
		},
		{
			Nums:     []int{0, 2, 4, 6, 8, 10},
			Expected: 1,
		},
		{
			Nums:     []int{},
			Expected: 0,
		},
		{
			Nums:     []int{1, 2, 0, 1},
			Expected: 3,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, longestConsecutive(tc.Nums), tc.Expected)
	}
}
