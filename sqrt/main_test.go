package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	assert.Equal(t, mySqrt(4), 2)
	assert.Equal(t, mySqrt(8), 2)
	assert.Equal(t, mySqrt(6), 2)
	assert.Equal(t, mySqrt(9), 3)
	assert.Equal(t, mySqrt(15), 3)
	assert.Equal(t, mySqrt(101), 10)
}
