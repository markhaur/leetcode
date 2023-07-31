package main

import "fmt"

func main() {
	fmt.Printf("firstBadVersion(%d): %d\n", 10, firstBadVersion(10))
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	low := 0
	badVersion := 1
	midVersion := 0
	isBad := false

	for low <= n {
		midVersion = low + (n-low)/2
		isBad = isBadVersion(midVersion) // this method is given by leetcode
		if isBad {
			badVersion = midVersion
			n = midVersion - 1
		} else {
			low = midVersion + 1
		}
	}

	return badVersion
}

// mocked
func isBadVersion(version int) bool {
	return true
}
