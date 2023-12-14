package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnesMinusZeros(t *testing.T) {
	tt := []struct {
		Grid     [][]int
		Expected [][]int
	}{
		{
			Grid:     [][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}},
			Expected: [][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}},
		},
		{
			Grid:     [][]int{{1, 1, 1}, {1, 1, 1}},
			Expected: [][]int{{5, 5, 5}, {5, 5, 5}},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, onesMinusZeros(tc.Expected))
	}
}
