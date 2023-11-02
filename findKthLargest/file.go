package main

import (
	"fmt"
	"os"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	if nums == nil {
		fmt.Println("Массив пуст")
		os.Exit(1)
	}
	if k > len(nums) || k < 1 {
		fmt.Println("k имеет некорректную величину")
		os.Exit(2)
	}

	n := len(nums)

	if k > n/2 {
		lowLimit := k - 1
		for i := n; i > lowLimit; i-- {
			for j := n - 1; j > n-i; j-- {
				if nums[j] < nums[j-1] {
					nums[j], nums[j-1] = nums[j-1], nums[j]
				}
			}
		}
		return nums[n-k]
	} else {
		lowLimit := n - k
		for i := n; i > lowLimit; i-- {
			for j := n - 1; j > n-i; j-- {
				if nums[j] > nums[j-1] {
					nums[j], nums[j-1] = nums[j-1], nums[j]
				}
			}
		}
		return nums[k-1]
	}
}

func main() {
	a := []int{10, 63, 55, 47, 84, 70, 41, 99, 6}

	for k := 1; k <= len(a); k++ {
		num := findKthLargest(a, k)
		fmt.Println(num)
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println(a)
}
