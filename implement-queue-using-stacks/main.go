package main

import "fmt"

func main() {
	queue := Constructor()

	queue.Push(1)
	queue.Push(8)
	fmt.Printf("pop: %d\n", queue.Pop())
	queue.Push(3)
	fmt.Printf("peek: %d\n", queue.Peek())
	fmt.Printf("pop: %d\n", queue.Pop())
	queue.Push(5)
	fmt.Printf("pop: %d\n", queue.Pop())
}
