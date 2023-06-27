package main

import "fmt"

func main() {
	fmt.Printf("hammingWeight of %d is %d\n", 11, hammingWeight(11))
	fmt.Printf("hammingWeight of %d is %d\n", 128, hammingWeight(128))
	fmt.Printf("hammingWeight of %d is %d\n", 4294967293, hammingWeight(4294967293))
}

/**
* Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).
 */
func hammingWeight(num uint32) int {
	counter := 0
	var remainder uint32
	for {
		remainder = num % 2
		if remainder == 1 {
			counter += 1
		}

		num = num / 2

		if num < 2 {
			if num == 1 {
				counter += 1
			}
			break
		}
	}

	return counter
}
