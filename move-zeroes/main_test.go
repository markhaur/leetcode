package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	tt := []struct {
		Nums     []int
		Expected []int
	}{
		{
			Nums:     []int{0, 1, 0, 2, 3, 4, 5},
			Expected: []int{1, 2, 3, 4, 5, 0, 0},
		},
		{
			Nums:     []int{0},
			Expected: []int{0},
		},
	}

	for _, tc := range tt {
		moveZeroes(tc.Nums)
		assert.True(t, matcharrays(tc.Nums, tc.Expected))
	}
}

func matcharrays(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
