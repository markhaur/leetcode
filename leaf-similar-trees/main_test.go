package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeafSimilar(t *testing.T) {
	root1 := generateTree1()
	root2 := generateTree2()

	assert.Equal(t, true, leafSimilar(root1, root2))
}
