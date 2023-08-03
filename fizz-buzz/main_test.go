package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	tt := []struct {
		N        int
		Expected []string
	}{
		{
			N:        3,
			Expected: []string{"1", "2", "Fizz"},
		},
		{
			N:        5,
			Expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			N:        15,
			Expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, tc := range tt {
		assert.True(t, match(fizzBuzz(tc.N), tc.Expected))
	}
}

func match(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
