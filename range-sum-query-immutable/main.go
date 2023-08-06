package main

import "fmt"

func main() {
	obj := Constructor([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("sum in range 0, 10 is %d", obj.SumRange(0, 9))
}

// https://leetcode.com/problems/range-sum-query-immutable/description
type NumArray struct {
	prefixArr []int
}

func Constructor(nums []int) NumArray {
	newArray := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		newArray[i+1] = nums[i] + newArray[i]
	}

	return NumArray{prefixArr: newArray}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixArr[right+1] - this.prefixArr[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
