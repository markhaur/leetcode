package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	tt := []struct {
		Nums     []int
		Expected []int
	}{
		{
			Nums:     []int{1, 2, 3, 4},
			Expected: []int{24, 12, 8, 6},
		},
		{
			Nums:     []int{-1, 1, 0, -3, 3},
			Expected: []int{0, 0, 9, 0, 0},
		},
		{
			Nums:     []int{-1, 1, 0, -3, 3, -1, 1, 0, -3, 3, -1, 1, 0, -3, 3, -1, 1, 0, -3, 3, -1, 1, 0, -3, 3, -1, 1, 0, -3, 3},
			Expected: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Nums:     []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Expected: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, productExceptSelf(tc.Nums), tc.Expected)
	}
}

/*
 * goos: windows
 * goarch: amd64
 * pkg: github.com/markhaur/leetcode/product-of-array-except-self
 * cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
 * BenchmarkProductExceptSelf-8   	50446294	        51.42 ns/op	       0 B/op	       0 allocs/o
 */
func BenchmarkProductExceptSelf(b *testing.B) {
	nums := []int{-1, 1, 0, -3, 3, -1, 1, 0, -3, 3, -1, 1, 0, -3, 3, -1, 1, 0, -3, 3, -1, 1, 0, -3, 3, -1, 1, 0, -3, 3}
	for i := 0; i < b.N; i++ {
		productExceptSelf(nums)
	}
}
