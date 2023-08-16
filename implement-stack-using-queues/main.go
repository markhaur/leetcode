package main

import "fmt"

func main() {
	stack := Constructor()
	stack.Push(5)
	stack.Push(7)

	fmt.Printf("Is Empty: %t\n", stack.Empty())
	fmt.Printf("Top: %d\n", stack.Top())
	fmt.Printf("Pop: %d\n", stack.Pop())
	fmt.Printf("Pop: %d\n", stack.Pop())
	fmt.Printf("Is Empty: %t\n", stack.Empty())
}
