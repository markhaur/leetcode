package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestAddDigits(t *testing.T) {
	assert.Equal(t, addDigits(38), 2)
	assert.Equal(t, addDigits(0), 0)
}
