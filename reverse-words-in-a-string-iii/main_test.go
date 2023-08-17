package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestReverseWords(t *testing.T) {
	tt := []struct {
		Word        string
		ReverseWord string
	}{
		{
			Word:        "Let's take LeetCode contest",
			ReverseWord: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			Word:        "God Ding",
			ReverseWord: "doG gniD",
		},
	}

	for _, tc := range tt {
		assert.Equal(t, reverseWords(tc.Word), tc.ReverseWord)
	}
}
