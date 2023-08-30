package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	tt := []struct {
		N        int
		Expected int
	}{
		{
			N:        2,
			Expected: 1,
		},
		{
			N:        3,
			Expected: 2,
		},
		{
			N:        4,
			Expected: 3,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, fib(tc.N), tc.Expected)
	}
}
