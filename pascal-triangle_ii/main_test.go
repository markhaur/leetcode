package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T) {
	tt := []struct {
		rowIndex int
		row      []int
	}{
		{
			rowIndex: 0,
			row:      []int{1},
		},
		{
			rowIndex: 1,
			row:      []int{1, 1},
		},
		{
			rowIndex: 2,
			row:      []int{1, 2, 1},
		},
		{
			rowIndex: 3,
			row:      []int{1, 3, 3, 1},
		},
		{
			rowIndex: 4,
			row:      []int{1, 4, 6, 4, 1},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, getRow(tc.rowIndex), tc.row)
	}
}
