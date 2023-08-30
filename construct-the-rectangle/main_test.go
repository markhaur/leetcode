package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConstructRectangle(t *testing.T) {
	tt := []struct {
		Area int
		ExpectedLW []int
	} {
		{
			Area: 4,
			ExpectedLW: []int{2,2},
		},
		{
			Area: 37,
			ExpectedLW: []int{37,1},
		},
		{
			Area: 122122,
			ExpectedLW: []int{427,286},
		},
		{
			Area: 100,
			ExpectedLW: []int{10,10},
		},
		{
			Area: 1000,
			ExpectedLW: []int{40,25},
		},
		{
			Area: 10001,
			ExpectedLW: []int{137,73},
		},
		{
			Area: 1,
			ExpectedLW: []int{1,1},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, constructRectangle(tc.Area), tc.ExpectedLW)
	}
}