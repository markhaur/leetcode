package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestTranspose(t *testing.T) {
	tt := []struct {
		Matrix   [][]int
		Expected [][]int
	}{
		{
			Matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			Expected: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			Matrix:   [][]int{{1, 2, 3}, {4, 5, 6}},
			Expected: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			Matrix:   [][]int{{3}},
			Expected: [][]int{{3}},
		},
	}

	for _, tc := range tt {
		assert.DeepEqual(t, tc.Expected, transpose(tc.Matrix))
	}
}
