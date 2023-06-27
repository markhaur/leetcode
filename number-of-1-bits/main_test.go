package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, hammingWeight(11), 3)
	assert.Equal(t, hammingWeight(128), 1)
	assert.Equal(t, hammingWeight(4294967293), 31)
}
