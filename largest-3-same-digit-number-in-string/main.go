package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "i6777133339"
	fmt.Printf("largest good integer: %s\n", largestGoodInteger(num))
}

func largestGoodInteger(num string) string {
	max := -1
	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			n, _ := strconv.Atoi(fmt.Sprintf("%c%c%c", num[i], num[i], num[i]))
			if n > max {
				max = n
			}
			i += 2
		}
	}

	if max == -1 {
		return ""
	}

	if max == 0 {
		return "000"
	}

	return strconv.Itoa(max)
}
