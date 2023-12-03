package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCharacters(t *testing.T) {
	tt := []struct {
		Words    []string
		Chars    string
		Expected int
	}{
		{
			Words:    []string{"cat", "bt", "hat", "tree"},
			Chars:    "atach",
			Expected: 6,
		},
		{
			Words:    []string{"hello", "world", "leetcode"},
			Chars:    "welldonehoneyr",
			Expected: 10,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, countCharacters(tc.Words, tc.Chars))
	}
}
