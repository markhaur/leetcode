package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeSumBST(t *testing.T) {
	root := generateTree()
	tt := []struct {
		Node     *TreeNode
		Low      int
		High     int
		Expected int
	}{
		{
			Node:     root,
			Low:      7,
			High:     15,
			Expected: 32,
		},
		{
			Node:     root,
			Low:      0,
			High:     20,
			Expected: 58,
		},
		{
			Node:     root,
			Low:      5,
			High:     10,
			Expected: 22,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, rangeSumBST(tc.Node, tc.Low, tc.High))
	}
}
