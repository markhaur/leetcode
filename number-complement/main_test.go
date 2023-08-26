package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindComplement(t *testing.T) {
	tt := []struct {
		N        int
		Expected int
	}{
		{
			N:        5,
			Expected: 2,
		},
		{
			N:        1,
			Expected: 0,
		},
		{
			N:        7,
			Expected: 0,
		},
		{
			N:        101,
			Expected: 26,
		},
		{
			N:        871568745,
			Expected: 202173078,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, findComplement(tc.N), tc.Expected)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/markhaur/leetcode/number-complement
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkFindComplement-8   	23574433	        82.78 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/markhaur/leetcode/number-complement	2.681s
*/
func BenchmarkFindComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findComplement(87156874598758956)
	}
}
