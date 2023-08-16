package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyStack(t *testing.T) {
	stack := Constructor()

	stack.Push(5)
	stack.Push(6)

	assert.NotNil(t, stack)
	assert.False(t, stack.Empty())
	assert.Equal(t, stack.Top(), 6)
	assert.Equal(t, stack.Pop(), 6)
	assert.Equal(t, stack.Pop(), 5)
	assert.True(t, stack.Empty())
}
