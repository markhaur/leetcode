package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Printf("Linked List has cycle: %t\n", hasCycle(generateLinkedListCycle()))
	fmt.Printf("Linked List has cycle: %t\n", hasCycle(generateLinkedListWithoutCycle()))
}

/**
* Given head, the head of a linked list, determine if the linked list has a cycle in it.
* There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. 
* Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
* Return true if there is a cycle in the linked list. Otherwise, return false.
* @link: https://leetcode.com/problems/linked-list-cycle/description/
*/
func hasCycle(head *ListNode) bool {
	iterator := head

	for iterator != nil {
		if iterator.Val > 100000 {
			return true
		}

		iterator.Val = 100001
		iterator = iterator.Next
	}

	return false
}

func generateLinkedListCycle() *ListNode {
	head := ListNode{Val: 1}

	first := ListNode{Val: 2}
	head.Next = &first

	second := ListNode{Val: 3}
	first.Next = &second

	third := ListNode{Val: 4}
	second.Next = &third

	fourth := ListNode{Val: 5}
	third.Next = &fourth

	fifth := ListNode{Val: 6}
	fourth.Next = &fifth

	fifth.Next = &head

	return &head
}

func generateLinkedListWithoutCycle() *ListNode {
	head := ListNode{Val: 1}

	first := ListNode{Val: 2}
	head.Next = &first

	second := ListNode{Val: 3}
	first.Next = &second

	third := ListNode{Val: 4}
	second.Next = &third

	fourth := ListNode{Val: 5}
	third.Next = &fourth

	fifth := ListNode{Val: 6}
	fourth.Next = &fifth

	return &head
}
