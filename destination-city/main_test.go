package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDestCity(t *testing.T) {
	tt := []struct {
		Paths    [][]string
		Expected string
	}{
		{
			Paths:    [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			Expected: "Sao Paulo",
		},
		{
			Paths:    [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			Expected: "A",
		},
		{
			Paths:    [][]string{{"A", "Z"}},
			Expected: "Z",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, destCity(tc.Paths))
	}
}
