package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t, climbStairs(2), 2)
	assert.Equal(t, climbStairs(3), 3)
	assert.Equal(t, climbStairs(4), 5)
	assert.Equal(t, climbStairs(5), 8)
	assert.Equal(t, climbStairs(6), 13)
}
