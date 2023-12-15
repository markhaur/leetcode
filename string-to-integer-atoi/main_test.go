package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	tt := []struct {
		S        string
		Expected int
	}{
		{
			S:        "42",
			Expected: 42,
		},
		{
			S:        "   -42",
			Expected: -42,
		},
		{
			S:        "4193 with words",
			Expected: 4193,
		},
		{
			S:        "-91283472332",
			Expected: -2147483648,
		},
		{
			S:        "+1",
			Expected: 1,
		},
		{
			S:        "21474836460",
			Expected: 2147483647,
		},
		{
			S:        "2147483648",
			Expected: 2147483647,
		},
		{
			S:        "9223372036854775808",
			Expected: 2147483647,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, myAtoi(tc.S))
	}
}
