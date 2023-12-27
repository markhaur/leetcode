package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCost(t *testing.T) {
	tt := []struct {
		Colors     string
		NeededTime []int
		Expected   int
	}{
		{
			Colors:     "abaac",
			NeededTime: []int{1, 2, 3, 4, 5},
			Expected:   3,
		},
		{
			Colors:     "abc",
			NeededTime: []int{1, 2, 3},
			Expected:   0,
		},
		{
			Colors:     "aabaa",
			NeededTime: []int{1, 2, 3, 4, 1},
			Expected:   2,
		},
		{
			Colors:     "bbbaaa",
			NeededTime: []int{4, 9, 3, 8, 8, 9},
			Expected:   23,
		},
		{
			Colors:     "aaabbbabbbb",
			NeededTime: []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1},
			Expected:   26,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, minCost(tc.Colors, tc.NeededTime))
	}
}
