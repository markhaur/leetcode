package main

import "fmt"

func main() {
	fmt.Printf("reverseBits(%d) => %d\n", 43261596, reverseBits(43261596))
	fmt.Printf("reverseBits(%d) => %d\n", 4294967293, reverseBits(4294967293))
}

func reverseBits(num uint32) uint32 {
	var bits [32]uint32
	for i := 0; i < 32; i++ {
		bits[i] = num % 2
		num = num / 2
		if num < 2 {
			bits[i+1] = num
			break
		}
	}

	pow := 31
	var reverseNumber uint32
	for _, bit := range bits {
		if bit == 1 {
			reverseNumber += power(2, pow)
		}
		pow -= 1
	}
	return reverseNumber
}

func power(base uint32, exponent int) uint32 {
	result := uint32(1)
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
