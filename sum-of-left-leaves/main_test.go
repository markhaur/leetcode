package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfLeftLeaves(t *testing.T) {
	root := generateBinaryTree()
	assert.Equal(t, sumOfLeftLeaves(root), 10)
}
