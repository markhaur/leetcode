package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tt := []struct {
		Strs     []string
		Expected [][]string
	}{
		{
			Strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			Expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			Strs:     []string{""},
			Expected: [][]string{{""}},
		},
		{
			Strs:     []string{"a"},
			Expected: [][]string{{"a"}},
		},
		{
			Strs:     []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"},
			Expected: [][]string{{"tin"}, {"pew"}, {"may"}, {"bar"}, {"cab"}, {"duh"}, {"ill"}, {"buy"}, {"max"}, {"doc"}},
		},
		{
			Strs:     []string{"mod", "mop", "pip", "tug", "hop", "dog", "met", "zoe", "axe", "mug", "fdr", "for", "fro", "fdr", "pap", "ray", "lop", "nth", "old", "eva", "ell", "mci", "wee", "ind", "but", "all", "her", "lew", "ken", "awl", "law", "rim", "zit", "did", "yam", "not", "ref", "lao", "gab", "sax", "cup", "new", "job", "new", "del", "gap", "win", "pot", "lam", "mgm", "yup", "hon", "khz", "sop", "has", "era", "ark"},
			Expected: [][]string{{"eva"}, {"ell"}, {"mgm"}, {"sop"}, {"her"}, {"yam"}, {"pot"}, {"khz"}, {"pip"}, {"wee"}, {"gap"}, {"has"}, {"yup"}, {"ray"}, {"lam"}, {"but"}, {"awl", "law"}, {"rim"}, {"new", "new"}, {"hon"}, {"tug"}, {"dog"}, {"nth"}, {"did"}, {"mod"}, {"era"}, {"mop"}, {"cup"}, {"not"}, {"zoe"}, {"mug"}, {"for", "fro"}, {"old"}, {"mci"}, {"job"}, {"ind"}, {"zit"}, {"win"}, {"ark"}, {"lao"}, {"gab"}, {"fdr", "fdr"}, {"lop"}, {"all"}, {"ken"}, {"ref"}, {"met"}, {"sax"}, {"axe"}, {"pap"}, {"lew"}, {"hop"}, {"del"}},
		},
	}

	for _, tc := range tt {
		assert.True(t, isMatch(groupAnagrams(tc.Strs), tc.Expected))
	}
}

/*
 * goos: windows
 * goarch: amd64
 * pkg: github.com/markhaur/leetcode/group-anagrams
 * cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
 * BenchmarkGroupAnagrams-8   	   48752	     40218 ns/op	   19447 B/op	      71 allocs/op
 */
func BenchmarkGroupAnagrams(b *testing.B) {
	strs := []string{"mod", "mop", "pip", "tug", "hop", "dog", "met", "zoe", "axe", "mug", "fdr", "for", "fro", "fdr", "pap", "ray", "lop", "nth", "old", "eva", "ell", "mci", "wee", "ind", "but", "all", "her", "lew", "ken", "awl", "law", "rim", "zit", "did", "yam", "not", "ref", "lao", "gab", "sax", "cup", "new", "job", "new", "del", "gap", "win", "pot", "lam", "mgm", "yup", "hon", "khz", "sop", "has", "era", "ark"}

	for i := 0; i < b.N; i++ {
		groupAnagrams(strs)
	}
}

func isMatch(str1, str2 [][]string) bool {
	if len(str1) != len(str2) {
		return false
	}

	for _, s1 := range str1 {
		found := false
		for _, s2 := range str2 {
			if reflect.DeepEqual(s1, s2) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
