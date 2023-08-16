package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyQueue(t *testing.T) {
	tt := []struct {
		OpsAndValidations func(t *testing.T, queue *MyQueue)
	}{
		{
			OpsAndValidations: func(t *testing.T, queue *MyQueue) {
				queue.Push(1)
				queue.Push(2)
				assert.Equal(t, queue.Peek(), 1)
				assert.Equal(t, queue.Pop(), 1)
				assert.False(t, queue.Empty())
			},
		},
		{
			OpsAndValidations: func(t *testing.T, queue *MyQueue) {
				queue.Push(1)
				queue.Push(2)
				assert.Equal(t, queue.Peek(), 1)
				assert.Equal(t, queue.Pop(), 1)
				assert.Equal(t, queue.Pop(), 2)
				assert.True(t, queue.Empty())
			},
		},
		{
			OpsAndValidations: func(t *testing.T, queue *MyQueue) {
				queue.Push(1)
				queue.Push(2)
				queue.Push(3)
				queue.Push(4)
				assert.Equal(t, queue.Pop(), 1)
				queue.Push(5)
				assert.Equal(t, queue.Pop(), 2)
				assert.Equal(t, queue.Pop(), 3)
				assert.Equal(t, queue.Pop(), 4)
				assert.Equal(t, queue.Pop(), 5)
				assert.True(t, queue.Empty())
			},
		},
		{
			OpsAndValidations: func(t *testing.T, queue *MyQueue) {
				queue.Push(1)
				assert.False(t, queue.Empty())
			},
		},
		{
			OpsAndValidations: func(t *testing.T, queue *MyQueue) {
				queue.Push(1)
				queue.Push(2)
				assert.Equal(t, queue.Pop(), 1)
				assert.Equal(t, queue.Peek(), 2)
				assert.False(t, queue.Empty())
			},
		},
		{
			OpsAndValidations: func(t *testing.T, queue *MyQueue) {
				queue.Push(1)
				queue.Push(2)
				queue.Push(3)
				assert.Equal(t, queue.Peek(), 1)
				assert.Equal(t, queue.Pop(), 1)
				assert.Equal(t, queue.Peek(), 2)
				assert.Equal(t, queue.Pop(), 2)
				assert.Equal(t, queue.Peek(), 3)
				assert.False(t, queue.Empty())
				assert.Equal(t, queue.Pop(), 3)
				assert.True(t, queue.Empty())
			},
		},
		{
			OpsAndValidations: func(t *testing.T, queue *MyQueue) {
				queue.Push(1)
				queue.Push(8)
				assert.Equal(t, queue.Pop(), 1)
				queue.Push(3)
				assert.Equal(t, queue.Peek(), 8)
				assert.Equal(t, queue.Pop(), 8)
				queue.Push(5)
				assert.Equal(t, queue.Pop(), 3)
				assert.False(t, queue.Empty())
			},
		},
	}

	for _, tc := range tt {
		queue := Constructor()
		assert.NotNil(t, queue)
		tc.OpsAndValidations(t, &queue)
	}
}
