package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumRange(t *testing.T) {
	tt := []struct {
		Nums   []int
		Left   int
		Right  int
		ExpSum int
	}{
		{
			Nums:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			Left:   0,
			Right:  9,
			ExpSum: 45,
		},
		{
			Nums:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			Left:   1,
			Right:  9,
			ExpSum: 45,
		},
		{
			Nums:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			Left:   5,
			Right:  9,
			ExpSum: 35,
		},
	}

	for _, tc := range tt {
		obj := Constructor(tc.Nums)
		assert.Equal(t, obj.SumRange(tc.Left, tc.Right), tc.ExpSum)
	}
}
