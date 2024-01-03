package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfBeams(t *testing.T) {
	tt := []struct {
		Bank     []string
		Expected int
	}{
		{
			Bank:     []string{"011001", "000000", "010100", "001000"},
			Expected: 8,
		},
		{
			Bank:     []string{"000", "111", "000"},
			Expected: 0,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, numberOfBeams(tc.Bank))
	}
}
