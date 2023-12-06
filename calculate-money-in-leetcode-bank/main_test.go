package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalMoney(t *testing.T) {
	tt := []struct {
		N        int
		Expected int
	}{
		{
			N:        4,
			Expected: 10,
		},
		{
			N:        10,
			Expected: 37,
		},
		{
			N:        20,
			Expected: 96,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, totalMoney(tc.N))
	}
}
