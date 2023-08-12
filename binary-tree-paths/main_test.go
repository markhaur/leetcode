package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreePaths(t *testing.T) {
	root := generateBinaryTree()
	expectedPaths := []string{"4->2->1", "4->2->3", "4->7->9"}
	actualPaths := binaryTreePaths(root)

	assert.Equal(t, len(actualPaths), len(expectedPaths))

	for i, _ := range actualPaths {
		assert.Equal(t, actualPaths[i], expectedPaths[i])
	}

}
