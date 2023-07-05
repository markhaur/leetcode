package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestReverseBits(t *testing.T) {
	assert.Equal(t, reverseBits(43261596), uint32(964176192))
	assert.Equal(t, reverseBits(4294967293), uint32(3221225471))
}
