package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	assert.Equal(t, majorityElement([]int{3, 2, 3}), 3)
	assert.Equal(t, majorityElement([]int{2,2,1,1,1,2,2}), 2)
	assert.Equal(t, majorityElement([]int{3,2,3,4,5,6,7,5,5,5,5,5,5,5}), 5)
}
