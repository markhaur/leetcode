package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T) {
	tt := []struct {
		Nums     []int
		Target   int
		Expected int
	}{
		{
			Nums:     []int{-1, 2, 1, -4},
			Target:   1,
			Expected: 2,
		},
		{
			Nums:     []int{0, 0, 0},
			Target:   1,
			Expected: 0,
		},
		{
			Nums:     []int{-1, 2, 1, -4},
			Target:   -1,
			Expected: -1,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, threeSumClosest(tc.Nums, tc.Target))
	}
}
