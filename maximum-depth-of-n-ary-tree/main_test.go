package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	root := generateNAryTree()
	assert.Equal(t, maxDepth(root), 5)
}
