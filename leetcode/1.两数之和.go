/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
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

// @lc code=end

// 暴力枚举
func twoSum1(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}