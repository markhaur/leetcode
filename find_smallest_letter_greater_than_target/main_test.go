package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreatestLetter(t *testing.T) {
	assert.Equal(t, nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'), byte('c'))
	assert.Equal(t, nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'), byte('f'))
	assert.Equal(t, nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z'), byte('x'))
}
