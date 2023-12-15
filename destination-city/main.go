package main

import "fmt"

func main() {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Printf("Destination City: %s\n", destCity(paths))

}

// https://leetcode.com/problems/destination-city/
func destCity(paths [][]string) string {
	out := make(map[string]int)

	for _, path := range paths {
		out[path[0]]++
	}

	for _, path := range paths {
		if _, ok := out[path[1]]; !ok {
			return path[1]
		}
	}

	return ""
}
