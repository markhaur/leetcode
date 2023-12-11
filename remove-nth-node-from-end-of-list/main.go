package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := generateList()
	head = removeNthFromEnd(head, 2)

	for head != nil {
		fmt.Printf("%d, ", head.Val)
		head = head.Next
	}
	fmt.Println("")
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	currentCounter := 0
	diffCounter := 0 - n
	currentNode := head
	deleteNode := head

	for currentNode.Next != nil {
		if diffCounter >= 0 {
			deleteNode = deleteNode.Next
		}
		currentNode = currentNode.Next
		diffCounter++
		currentCounter++
	}
	if currentCounter == n-1 {
		head = head.Next
		return head
	}
	deleteNode.Next = deleteNode.Next.Next
	return head
}

func generateList() *ListNode {
	five := ListNode{Val: 5}
	four := ListNode{Val: 4, Next: &five}
	three := ListNode{Val: 3, Next: &four}
	two := ListNode{Val: 2, Next: &three}
	head := ListNode{Val: 1, Next: &two}
	return &head
}
