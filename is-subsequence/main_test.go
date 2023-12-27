package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	tt := []struct {
		S        string
		T        string
		Expected bool
	}{
		{
			S:        "abc",
			T:        "ahbgdc",
			Expected: true,
		},
		{
			S:        "axc",
			T:        "ahbgdc",
			Expected: false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, isSubsequence(tc.S, tc.T))
	}
}
