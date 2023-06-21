package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDepth(t *testing.T) {
	root := generateTree()
	assert.Equal(t, minDepth(root), 3)

	rightTree := generateRightTree()
	assert.Equal(t, minDepth(rightTree), 5)
}
