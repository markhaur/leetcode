package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	tree := generateTree()
	assert.True(t, isSameTree(tree, tree))

	otherTree := generateTree()
	otherTree.Val = 2
	assert.False(t, isSameTree(tree, otherTree))
}
