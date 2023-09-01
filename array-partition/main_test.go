package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	tt := []struct {
		Nums     []int
		Expected int
	}{
		{
			Nums:     []int{1, 4, 3, 2},
			Expected: 4,
		},
		{
			Nums:     []int{6, 2, 6, 5, 1, 2},
			Expected: 9,
		},
		{
			Nums:     []int{6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2},
			Expected: 88,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, arrayPairSum(tc.Nums), tc.Expected)
	}
}
