package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNegative(t *testing.T) {
	tt := []struct {
		grid          [][]int
		expectedCount int
	}{
		{
			grid: [][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			},
			expectedCount: 8,
		},
		{
			grid: [][]int{
				{4, 3, 2, -1, -2},
				{3, 2, 1, -1, -5},
				{1, 1, -1, -2, -3},
				{-1, -1, -2, -3, -9},
			},
			expectedCount: 12,
		},
		{
			grid: [][]int{
				{4},
				{3},
				{1},
				{1},
			},
			expectedCount: 0,
		},
		{
			grid:          [][]int{},
			expectedCount: 0,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, countNegatives(tc.grid), tc.expectedCount)
	}
}
