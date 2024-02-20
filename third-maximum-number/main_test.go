package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	tt := []struct {
		Nums     []int
		Expected int
	}{
		{
			Nums:     []int{3, 2, 1},
			Expected: 1,
		},
		{
			Nums:     []int{1, 2},
			Expected: 2,
		},
		{
			Nums:     []int{2, 2, 3, 1},
			Expected: 1,
		},
		{
			Nums:     []int{1, 1, 2},
			Expected: 2,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, thirdMax(tc.Nums))
	}
}
