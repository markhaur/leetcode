package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyChoco(t *testing.T) {
	tt := []struct {
		Prices   []int
		Money    int
		Expected int
	}{
		{
			Prices:   []int{1, 2, 2},
			Money:    3,
			Expected: 0,
		},
		{
			Prices:   []int{3, 2, 3},
			Money:    3,
			Expected: 3,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, buyChoco(tc.Prices, tc.Money))
	}
}
