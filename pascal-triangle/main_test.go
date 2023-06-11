package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	tt := []struct {
		numRows  int
		triangle [][]int
	}{
		{
			numRows:  1,
			triangle: [][]int{{1}},
		},
		{
			numRows:  2,
			triangle: [][]int{{1}, {1, 1}},
		},
		{
			numRows:  3,
			triangle: [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			numRows:  4,
			triangle: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
		{
			numRows:  5,
			triangle: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, generate(tc.numRows), tc.triangle)
	}
}
