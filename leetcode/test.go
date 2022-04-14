package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, x := range nums {
		if p, ok := hashmap[target-x]; ok {
			return []int{p, i}
		}
		hashmap[x] = i
	}
	return nil
}
func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
