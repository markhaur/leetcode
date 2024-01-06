package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobScheduling(t *testing.T) {
	tt := []struct {
		StartTime []int
		EndTime   []int
		Profit    []int
		Expected  int
	}{
		{
			StartTime: []int{1, 2, 3, 3},
			EndTime:   []int{3, 4, 5, 6},
			Profit:    []int{50, 10, 40, 70},
			Expected:  120,
		},
		{
			StartTime: []int{1, 2, 3, 4, 6},
			EndTime:   []int{3, 5, 10, 6, 9},
			Profit:    []int{20, 20, 100, 70, 60},
			Expected:  150,
		},
		{
			StartTime: []int{1, 1, 1},
			EndTime:   []int{2, 3, 4},
			Profit:    []int{5, 6, 4},
			Expected:  6,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, jobScheduling(tc.StartTime, tc.EndTime, tc.Profit))
	}
}
