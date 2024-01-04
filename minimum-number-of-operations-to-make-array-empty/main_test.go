package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	tt := []struct {
		Nums     []int
		Expected int
	}{
		{
			Nums:     []int{2, 3, 3, 2, 2, 4, 2, 3, 4},
			Expected: 4,
		},
		{
			Nums:     []int{2,1,2,2,3,3},
			Expected: -1,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, minOperations(tc.Nums))
	}
}