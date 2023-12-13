package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tt := []struct {
		Nums     []int
		Target   int
		Expected int
	}{
		{
			Nums:     []int{4, 5, 6, 7, 0, 1, 2},
			Target:   0,
			Expected: 4,
		},
		{
			Nums:     []int{4, 5, 6, 7, 0, 1, 2},
			Target:   3,
			Expected: -1,
		},
		{
			Nums:     []int{1},
			Target:   0,
			Expected: -1,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, search(tc.Nums, tc.Target))
	}
}
