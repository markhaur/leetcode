package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestMissingNumber(t *testing.T) {
	tt := []struct {
		Nums          []int
		MissingNumber int
	}{
		{
			Nums:          []int{3, 0, 1},
			MissingNumber: 2,
		},
		{
			Nums:          []int{0, 1},
			MissingNumber: 2,
		},
		{
			Nums:          []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			MissingNumber: 8,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, missingNumber(tc.Nums), tc.MissingNumber)
	}
}
