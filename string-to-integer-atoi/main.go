package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "9223372036854775808 with words"
	fmt.Printf("integer value with myAtoi: %d\n", myAtoi(s))
}

// https://leetcode.com/problems/string-to-integer-atoi/
func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}

	numericMap := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	i := 0
	isminus := false
	if s[0] == '-' {
		isminus = true
		i = 1
	}
	if s[0] == '+' {
		i = 1
	}

	number := 0
	for i = i; i < len(s); i++ {
		num, found := numericMap[s[i]]
		if !found {
			break
		}
		number = number*10 + num
		if isminus {
			if -1*number < -2147483648 {
				return -2147483648
			}
		} else if number >= 2147483648 {
			return 2147483647
		}
	}

	if isminus {
		return -1 * number
	}
	return number
}
