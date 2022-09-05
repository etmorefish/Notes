package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7}
	fmt.Println(append(a, b...))
	fmt.Println(QuickSort([]int{4, 2, 7, 4, 8, 5, 9, 1, 0}))
}

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	low := make([]int, 0, 0)
	high := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	flag := nums[0]
	mid = append(mid, flag)
	for i := 1; i < len(nums); i++ {
		if nums[i] < flag {
			low = append(low, nums[i])
		} else if nums[i] > flag {
			high = append(high, nums[i])
		} else {
			mid = append(mid, nums[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	return append(append(low, mid...), high...)
}
