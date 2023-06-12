package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPathSum(t *testing.T) {
	root := generateBinaryTree()

	assert.True(t, hasPathSum(root, 22))
	assert.True(t, hasPathSum(root, 27))
	assert.True(t, hasPathSum(root, 26))
	assert.True(t, hasPathSum(root, 18))
	assert.False(t, hasPathSum(root, 19))
}
