package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContentChildren(t *testing.T) {
	tt := []struct {
		G        []int
		S        []int
		Expected int
	}{
		{
			G:        []int{1, 2, 3},
			S:        []int{1, 1},
			Expected: 1,
		},
		{
			G:        []int{1, 2},
			S:        []int{1, 2, 3},
			Expected: 2,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, findContentChildren(tc.G, tc.S))
	}
}
