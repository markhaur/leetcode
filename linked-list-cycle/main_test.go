package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	assert.True(t, hasCycle(generateLinkedListCycle()))
	assert.False(t, hasCycle(generateLinkedListWithoutCycle()))
}
