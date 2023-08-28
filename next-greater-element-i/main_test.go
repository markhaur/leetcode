package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {
	tt := []struct {
		Nums1    []int
		Nums2    []int
		Expected []int
	}{
		{
			Nums1:    []int{4, 1, 2},
			Nums2:    []int{1, 3, 4, 2},
			Expected: []int{-1, 3, -1},
		},
		{
			Nums1:    []int{2, 4},
			Nums2:    []int{1, 2, 3, 4},
			Expected: []int{3, -1},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, nextGreaterElement(tc.Nums1, tc.Nums2), tc.Expected)
	}
}

/*
 * goos: windows
 * goarch: amd64
 * pkg: github.com/markhaur/leetcode/next-greater-element-i
 * cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
 * BenchmarkNextGreaterElement-8   	 6409407	       189.6 ns/op	      24 B/op	       1 allocs/op
 */
func BenchmarkNextGreaterElement(b *testing.B) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	for i := 0; i < b.N; i++ {
		nextGreaterElement(nums1, nums2)
	}
}
