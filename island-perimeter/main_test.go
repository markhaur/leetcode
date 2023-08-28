package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslandPerimeter(t *testing.T) {
	tt := []struct {
		Grid      [][]int
		Perimeter int
	}{
		{
			Grid:      [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}},
			Perimeter: 16,
		},
		{
			Grid:      [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}, {0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}},
			Perimeter: 30,
		},
		{
			Grid:      [][]int{{1}},
			Perimeter: 4,
		},
		{
			Grid:      [][]int{{1, 0}},
			Perimeter: 4,
		},
		{
			Grid:      [][]int{{0}, {1}},
			Perimeter: 4,
		},
		{
			Grid:      [][]int{{0, 1, 1}},
			Perimeter: 6,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, islandPerimeter(tc.Grid), tc.Perimeter)
	}
}

/*
 * goos: windows
 * goarch: amd64
 * pkg: github.com/markhaur/leetcode/island-perimeter
 * cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
 * BenchmarkIslandPerimeter-8   	10738696	        95.15 ns/op	       0 B/op	       0 allocs/op
 */
func BenchmarkIslandPerimeter(b *testing.B) {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}, {0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}

	for i := 0; i < b.N; i++ {
		islandPerimeter(grid)
	}

}
