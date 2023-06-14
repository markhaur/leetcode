package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMinimumDifference(t *testing.T) {
	assert.Equal(t, getMinimumDifference(generateBST()), 1)
}
