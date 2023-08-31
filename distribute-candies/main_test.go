package main

import (
	"testing"
	"github.com/stretchr/testify/assert"

)
func TestDistributeCandies(t *testing.T) {
	tt := []struct {
		CandyType []int
		Expected int
	} {
		{
			CandyType: []int{1,1,2,2,3,3},
			Expected: 3,
		},
		{
			CandyType: []int{1,1,2,3},
			Expected: 2,
		},
		{
			CandyType: []int{6,6,6,6},
			Expected: 1,
		},
		{
			CandyType: []int{1,1,2,2,3,3,9,3,4,6,7,8,11,12},
			Expected: 7,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, distributeCandies(tc.CandyType), tc.Expected)
	}
}