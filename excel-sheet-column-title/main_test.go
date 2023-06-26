package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestConvertToTitle(t *testing.T) {
	assert.Equal(t, convertToTitle(1), "A")
	assert.Equal(t, convertToTitle(28), "AB")
	assert.Equal(t, convertToTitle(701), "ZY")
	assert.Equal(t, convertToTitle(52), "AZ")
	assert.Equal(t, convertToTitle(78), "BZ")
	assert.Equal(t, convertToTitle(53), "BA")
	assert.Equal(t, convertToTitle(702), "ZZ")
}
