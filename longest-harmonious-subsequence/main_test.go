package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLHS(t *testing.T) {
	tt := []struct {
		Nums     []int
		Expected int
	}{
		{
			Nums:     []int{1, 3, 2, 2, 5, 2, 3, 7},
			Expected: 5,
		},
		{
			Nums:     []int{1, 2, 3, 4},
			Expected: 2,
		},
		{
			Nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			Expected: 0,
		},
		{
			Nums:     []int{1, 2, 2, 5, 2, 7},
			Expected: 4,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, findLHS(tc.Nums))
	}
}
