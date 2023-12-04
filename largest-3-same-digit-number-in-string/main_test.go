package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestGoodInteger(t *testing.T) {
	tt := []struct {
		Num      string
		Expected string
	}{
		{
			Num:      "2300019",
			Expected: "000",
		},
		{
			Num:      "i6777133339",
			Expected: "777",
		},
		{
			Num:      "42352338",
			Expected: "",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, largestGoodInteger(tc.Num))
	}
}
