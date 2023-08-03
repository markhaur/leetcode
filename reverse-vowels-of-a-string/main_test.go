package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	assert.Equal(t, reverseVowels(""), "")
	assert.Equal(t, reverseVowels("hello"), "holle")
	assert.Equal(t, reverseVowels("leetcode"), "leotcede")
	assert.Equal(t, reverseVowels("aeiou"), "uoiea")
	assert.Equal(t, reverseVowels("llllllll"), "llllllll")
	assert.Equal(t, reverseVowels("aA"), "Aa")
}
