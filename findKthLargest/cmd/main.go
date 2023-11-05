package main

import (
	"fmt"
	"os"
	"sort"
)

// func partialSort: сортирует k чисел и перемещает в начало массива
func partialSort(a []int, k int, less func(e1, e2 int) bool) {
	n := len(a)
	for i := n; i > k; i-- {
		for j := n - 1; j > n-i; j-- {
			if less(a[j], a[j-1]) {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

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
	numsToSort := n - k

	if k > n/2 {
		numsToSort = k - 1
		partialSort(nums, numsToSort, func(e1, e2 int) bool {
			return e1 < e2
		})
		return nums[n-k]
	} else {
		partialSort(nums, numsToSort, func(e1, e2 int) bool {
			return e1 > e2
		})
		return nums[k-1]
	}
}

func test(a []int, k int) {
	fmt.Println("nums:", a, "k =", k)
	num := findKthLargest(a, k)
	fmt.Println("result =", num)

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println("Sorted array:", a)
	fmt.Println()
}

func main() {
	test([]int{10, 63, 55, 47, 84, 70, 41, 99, 6}, 2)

	test([]int{44, 7, 15, 84, 83, 73, 29, 77, 2}, 5)

	test([]int{42, 48, 19, 99, 50, 90, 80, 24, 10}, 8)
}
