package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfThree(t *testing.T) {
	assert.True(t, isPowerOfThree(3))
	assert.True(t, isPowerOfThree(1))
	assert.True(t, isPowerOfThree(81))
	assert.False(t, isPowerOfThree(0))
	assert.False(t, isPowerOfThree(-1))
}
