package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindDisappearedNumbers(t *testing.T) {
	tt := []struct {
		Nums     []int
		Expected []int
	}{
		{
			Nums:     []int{1, 2, 3, 4, 4, 2, 3, 6},
			Expected: []int{5, 7, 8},
		},
		{
			Nums:     []int{1, 1, 1, 1},
			Expected: []int{2, 3, 4},
		},
		{
			Nums:     []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, isEqual(findDisappearedNumbers(tc.Nums), tc.Expected), true)
	}
}

func isEqual(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}

	for i, _ := range nums1 {
		if nums1[i] != nums2[i] {
			return false
		}
	}

	return true
}
