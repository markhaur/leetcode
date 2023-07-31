package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	tt := []struct {
		Nums1    []int
		Nums2    []int
		Expected []int
	}{
		{
			Nums1:    []int{1, 2, 2, 1},
			Nums2:    []int{2, 2},
			Expected: []int{2},
		},
		{
			Nums1:    []int{4, 9, 5},
			Nums2:    []int{9, 4, 9, 8, 4},
			Expected: []int{9, 4},
		},
	}

	for _, tc := range tt {
		assert.True(t, match(intersection(tc.Nums1, tc.Nums2), tc.Expected))
	}
}

func match(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for _, a1 := range arr1 {
		found := false
		for _, a2 := range arr2 {
			if a1 == a2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
