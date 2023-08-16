package main

// https://leetcode.com/problems/implement-queue-using-stacks/description
type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) > 0 {
		val := this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[:len(this.outStack)-1]
		return val
	}
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
	return this.Pop()
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) > 0 {
		return this.outStack[len(this.outStack)-1]
	}
	if len(this.inStack) > 0 {
		return this.inStack[0]
	}
	return -1
}

func (this *MyQueue) Empty() bool {
	return len(this.outStack) == 0 && len(this.inStack) == 0
}
