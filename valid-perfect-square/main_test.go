package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	assert.True(t, isPerfectSquare(4))
	assert.True(t, isPerfectSquare(16))
	assert.True(t, isPerfectSquare(1))
	assert.False(t, isPerfectSquare(0))
	assert.False(t, isPerfectSquare(-1))
}
