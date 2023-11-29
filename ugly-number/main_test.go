package main

import (
	"math"
	"testing"

	"gotest.tools/assert"
)

func TestIsUgly(t *testing.T) {
	tt := []struct {
		Num      int
		Expected bool
	}{
		{
			Num:      1,
			Expected: true,
		},
		{
			Num:      0,
			Expected: false,
		},
		{
			Num:      55,
			Expected: false,
		},
		{
			Num:      14,
			Expected: false,
		},
		{
			Num:      2,
			Expected: true,
		},
		{
			Num:      8,
			Expected: true,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.Expected, isUgly(tc.Num))
	}
}

/*
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkIsUgly
BenchmarkIsUgly-8
100000000               10.98 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/markhaur/ugly-number 1.730s
*/
func BenchmarkIsUgly(b *testing.B) {
	num := int(math.MaxInt64)

	for i := 0; i < b.N; i++ {
		isUgly(num)
	}

}
