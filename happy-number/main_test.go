package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	assert.True(t, isHappy(19))
	assert.True(t, isHappy(1))
	assert.False(t, isHappy(2))
	assert.False(t, isHappy(18))
	assert.False(t, isHappy(125))
	assert.False(t, isHappy(2147483647))
}
