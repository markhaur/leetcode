package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		Numbers  []int
		Target   int
		Expected []int
	}{
		{
			Numbers:  []int{2, 7, 11, 15},
			Target:   9,
			Expected: []int{1, 2},
		},
		{
			Numbers:  []int{2, 3, 4},
			Target:   6,
			Expected: []int{1, 3},
		},
		{
			Numbers:  []int{-1, 0},
			Target:   -1,
			Expected: []int{1, 2},
		},
		{
			Numbers:  []int{2, 7, 11, 15},
			Target:   26,
			Expected: []int{3, 4},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, twoSum(tc.Numbers, tc.Target))
	}
}
