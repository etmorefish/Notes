/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */
package leetcode
// @lc code=start
func search(nums []int, target int) int {
    low, high := 0, len(nums) -1
    for low <= high{
        mid := (high - low)/2 + low
        num := nums[mid]
        if num == target{
            return mid
        }else if num > target{
            high = mid - 1
        }else{
            low = mid + 1
        }
    }
    return -1
}
// @lc code=end

