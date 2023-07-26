package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome(generatePalindromeLinkedList()))
	assert.False(t, isPalindrome(generateNonPalindromeLinkedList()))
}
