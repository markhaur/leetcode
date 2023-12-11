package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSpecialInteger(t *testing.T) {
	tt := []struct {
		Arr      []int
		Expected int
	}{
		{
			Arr:      []int{1, 2, 2, 6, 6, 6, 6, 7, 10},
			Expected: 6,
		},
		{
			Arr:      []int{1, 1},
			Expected: 1,
		},
		{
			Arr:      []int{1, 2, 3, 3},
			Expected: 3,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, findSpecialInteger(tc.Arr))
	}
}
