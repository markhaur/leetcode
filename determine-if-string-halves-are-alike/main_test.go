package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHalvesAreAlike(t *testing.T) {
	tt := []struct {
		S        string
		Expected bool
	}{
		{
			S:        "book",
			Expected: true,
		},
		{
			S:        "textbook",
			Expected: false,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, halvesAreAlike(tc.S))
	}
}
