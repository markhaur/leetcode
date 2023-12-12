package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
	tt := []struct {
		Nums     []int
		Expected int
	}{
		{
			Nums:     []int{3, 4, 5, 2},
			Expected: 12,
		},
		{
			Nums:     []int{1, 5, 4, 5},
			Expected: 16,
		},
		{
			Nums:     []int{3, 7},
			Expected: 12,
		},
		{
			Nums:     []int{2, 2, 1, 8, 1, 5, 4, 5, 2, 10, 3, 6, 5, 2, 3},
			Expected: 63,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, maxProduct(tc.Nums))
	}
}
