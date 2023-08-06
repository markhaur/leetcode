package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNodes(t *testing.T) {
	assert.Equal(t, countNodes(generateBinaryTree()), 6)
}
