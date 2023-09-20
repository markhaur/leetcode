package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tt := []struct {
		Nums     []int
		Expected [][]int
	}{
		{
			Nums:     []int{-1, 0, 1, 2, -1, -4},
			Expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			Nums:     []int{0, 1, 1},
			Expected: nil,
		},
		{
			Nums:     []int{0, 0, 0},
			Expected: [][]int{{0, 0, 0}},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, threeSum(tc.Nums))
	}
}
