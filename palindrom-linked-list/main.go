package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	palindromeLinkedList := generatePalindromeLinkedList()
	nonPalindromeLinkedList := generateNonPalindromeLinkedList()

	fmt.Printf("is linkedlist palindrome: %t\n", isPalindrome(nonPalindromeLinkedList))
	fmt.Printf("is linkedlist palindrome: %t\n", isPalindrome(palindromeLinkedList))

}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// Find the middle of the linked list
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half of the linked list
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// Compare the first and second half of the linked list
	for prev != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}

	return true
}

func generatePalindromeLinkedList() *ListNode {
	head := ListNode{Val: 1}
	first := ListNode{Val: 2}
	second := ListNode{Val: 2}
	third := ListNode{Val: 1}

	head.Next = &first
	first.Next = &second
	second.Next = &third

	return &head
}

func generateNonPalindromeLinkedList() *ListNode {
	head := ListNode{Val: 1}
	first := ListNode{Val: 2}
	second := ListNode{Val: 3}
	third := ListNode{Val: 1}

	head.Next = &first
	first.Next = &second
	second.Next = &third

	return &head
}
