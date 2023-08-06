package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	assert.True(t, isAnagram("nagaram", "anagram"))
	assert.True(t, isAnagram("listen", "silent"))
	assert.True(t, isAnagram("evil", "vile"))
	assert.True(t, isAnagram("debit", "bidet"))
	assert.True(t, isAnagram("desserts", "stressed"))
	assert.False(t, isAnagram("bb", "ac"))
	assert.False(t, isAnagram("car", "rat"))
}
