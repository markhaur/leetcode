package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	tt := []struct {
		Words    []string
		Expected []string
	}{
		{
			Words:    []string{"Hello", "Alaska", "Dad", "Peace"},
			Expected: []string{"Alaska", "Dad"},
		},
		{
			Words:    []string{"omk"},
			Expected: []string{},
		},
		{
			Words:    []string{"adsdf", "sfd"},
			Expected: []string{"adsdf", "sfd"},
		},
	}

	for _, tc := range tt {
		assert.True(t, match(findWords(tc.Words), tc.Expected))
	}
}

/*
*
goos: windows
goarch: amd64
pkg: github.com/markhaur/leetcode/keyboard-row
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkFindWords-8   	 4067043	       624.6 ns/op	      48 B/op	       2 allocs/op
PASS
ok  	github.com/markhaur/leetcode/keyboard-row	3.359s
*/
func BenchmarkFindWords(b *testing.B) {
	words := []string{"Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad",
		"Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace",
		"Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello",
		"Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska",
		"Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace",
		"Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska",
		"Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello",
		"Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace",
		"Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad",
		"Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska",
		"Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace", "Hello",
		"Alaska", "Dad", "Peace", "Hello", "Alaska", "Dad", "Peace"}
	for i := 0; i < b.N; i++ {
		findWords(words)
	}
}

func match(words1, words2 []string) bool {
	if len(words1) != len(words2) {
		return false
	}

	for _, w1 := range words1 {
		found := false
		for _, w2 := range words2 {
			if w1 == w2 {
				found = true
				break
			}
		}
		if found == false {
			return false
		}
	}
	return true
}
