package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummaryRanges(t *testing.T) {
	tt := []struct {
		nums   []int
		ranges []string
	}{
		{
			nums:   []int{0, 1, 2, 4, 5, 7},
			ranges: []string{"0->2", "4->5", "7"},
		},
		{
			nums:   []int{0, 2, 3, 4, 6, 8, 9},
			ranges: []string{"0", "2->4", "6", "8->9"},
		},
		{
			nums:   []int{},
			ranges: nil,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, summaryRanges(tc.nums), tc.ranges)
	}
}
