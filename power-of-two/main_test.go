package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	assert.True(t, isPowerOfTwo(1))
	assert.True(t, isPowerOfTwo(16))
	assert.True(t, isPowerOfTwo(32))

	assert.False(t, isPowerOfTwo(3))
	assert.False(t, isPowerOfTwo(15))
	assert.False(t, isPowerOfTwo(100))
}
