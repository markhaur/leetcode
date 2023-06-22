package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBalanced(t *testing.T) {
	assert.True(t, isBalanced(generateBalancedTree()))
	assert.True(t, isBalanced(nil))
	assert.False(t, isBalanced(generateUnbalancedTree()))
}
