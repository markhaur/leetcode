package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	tt := []struct {
		Nums     []int
		K        int
		Expected []int
	}{
		{
			Nums:     []int{1, 1, 1, 2, 2, 3},
			K:        2,
			Expected: []int{1, 2},
		},
		{
			Nums:     []int{1},
			K:        1,
			Expected: []int{1},
		},
		{
			Nums:     []int{1, 1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8},
			K:        8,
			Expected: []int{7, 8, 4, 5, 6, 1, 2, 3},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, topKFrequent(tc.Nums, tc.K), tc.Expected)
	}
}
