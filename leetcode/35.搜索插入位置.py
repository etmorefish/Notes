#
# @lc app=leetcode.cn id=35 lang=python3
#
# [35] 搜索插入位置
#  二分查找

# @lc code=start
from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) -1
        while left <= right:
            mid = (right - left)//2 + left
            num = nums[mid]
            if num > target:
                right = mid -1
            elif num < target:
                left = mid +1
            else:
                return mid
        return right +1
                
# @lc code=end

