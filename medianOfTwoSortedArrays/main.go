package main

func main() {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mid := len(nums1) + len(nums2)
	prev := -1
	if mid%2 == 0 {
		mid = mid / 2
		prev = mid - 1
	}

	return -1
}
