package main

func findKthLargest(nums []int, k int) int {
	if nums == nil {
		panic("Массив пуст")
	}
	if k > len(nums) {
		panic("k больше длины массива")
	}
	n := len(nums) - 1
	for i := n; i > k; i-- {
		for j := n; j > n-i; j-- {
			if nums[j] > nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}

	return nums[k-1]
}
