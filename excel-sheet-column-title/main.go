package main

import "fmt"

func main() {
	fmt.Printf("Excel column sheet title for %d is %s\n", 1, convertToTitle(1))
	fmt.Printf("Excel column sheet title for %d is %s\n", 28, convertToTitle(28))
	fmt.Printf("Excel column sheet title for %d is %s\n", 701, convertToTitle(701))
	fmt.Printf("Excel column sheet title for %d is %s\n", 52, convertToTitle(52))
	fmt.Printf("Excel column sheet title for %d is %s\n", 78, convertToTitle(78))
	fmt.Printf("Excel column sheet title for %d is %s\n", 53, convertToTitle(53))
	fmt.Printf("Excel column sheet title for %d is %s\n", 702, convertToTitle(702))
}

/**
* Given an integer `columnNumber`, return its corresponding column title as it appears in an Excel sheet.
* For example
* A -> 1
* B -> 2
* C -> 3
* ...
* Z -> 26
* AA -> 27
* AB -> 28
 */
func convertToTitle(columnNumber int) string {
	letterMap := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	remainder := 0
	columnTitle := ""
	for {
		if columnNumber <= 26 {
			columnTitle = fmt.Sprintf("%c%s", letterMap[columnNumber-1], columnTitle)
			break
		}

		remainder = columnNumber % 26
		if remainder == 0 {
			columnTitle = fmt.Sprintf("%c%s", 'Z', columnTitle)
			columnNumber = (columnNumber / 26) - 1
			continue
		}

		columnNumber = columnNumber / 26
		columnTitle = fmt.Sprintf("%c%s", letterMap[remainder-1], columnTitle)
	}

	return columnTitle
}
