package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
	root := generateTree()
	assert.True(t, isSymmetric(root))

	// making tree asymmetric
	root.Right.Val = 3
	assert.False(t, isSymmetric(root))
}
