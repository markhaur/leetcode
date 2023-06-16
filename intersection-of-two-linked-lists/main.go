package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	headA, headB, intersectNode := generateListOfNodes()
	fmt.Printf("getIntersectionNode function found correct intersection node: %t\n", getIntersectionNode(headA, headB) == intersectNode)
}

// https://leetcode.com/problems/intersection-of-two-linked-lists/description/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	iteratorA, iteratorB := headA, headB
	var foundNode *ListNode

	for iteratorA != nil || iteratorB != nil {
		if iteratorA == iteratorB {
			foundNode = iteratorA
			break
		}

		if iteratorA != nil {
			if iteratorA.Val < 0 {
				foundNode = iteratorA
				break
			}
			iteratorA.Val = iteratorA.Val * -1
			iteratorA = iteratorA.Next
		}

		if iteratorB != nil {
			if iteratorB.Val < 0 {
				foundNode = iteratorB
				break
			}
			iteratorB.Val = iteratorB.Val * -1
			iteratorB = iteratorB.Next
		}
	}

	chA := resetList(headA)
	chB := resetList(headB)

	<-chA
	<-chB
	return foundNode
}

func resetList(head *ListNode) <-chan bool {
	ch := make(chan bool)
	go func() {
		for head != nil && head.Val < 0 {
			head.Val = head.Val * -1
			head = head.Next
		}
		ch <- true
	}()
	return ch
}

func generateListOfNodes() (headA, headB, intersectNode *ListNode) {
	common3 := ListNode{Val: 5}
	common2 := ListNode{Val: 4, Next: &common3}
	common1 := ListNode{Val: 8, Next: &common2}

	aNode1 := ListNode{Val: 1, Next: &common1}
	aNode2 := ListNode{Val: 4, Next: &aNode1}

	bNode1 := ListNode{Val: 1, Next: &common1}
	bNode2 := ListNode{Val: 6, Next: &bNode1}
	bNode3 := ListNode{Val: 5, Next: &bNode2}

	return &aNode2, &bNode3, &common1
}
