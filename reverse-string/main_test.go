package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	tt := []struct {
		ByteArray []byte
		Expected  []byte
	}{
		{
			ByteArray: []byte{'h', 'e', 'l', 'l', 'o'},
			Expected:  []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			ByteArray: []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			Expected:  []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
		{
			ByteArray: []byte{'H'},
			Expected:  []byte{'H'},
		},
		{
			ByteArray: []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o'},
			Expected:  []byte{'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a', '9', '8', '7', '6', '5', '4', '3', '2', '1'},
		},
	}

	for _, tc := range tt {
		reverseString(tc.ByteArray)
		assert.True(t, match(tc.ByteArray, tc.Expected))
	}
}

func match(arr1, arr2 []byte) bool {
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
