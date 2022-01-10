#
# @lc app=leetcode.cn id=977 lang=python3
#
# [977] 有序数组的平方
# 双指针

# @lc code=start
from typing import List


class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        # return sorted(num * num for num in nums)
        n = len(nums)
        ans = [0] * n
        
        i, j, pos = 0, n-1, n-1
        while i <= j:
            if nums[i] * nums[i] > nums[j] * nums[j]:
                ans[pos] = nums[i] * nums[i]
                i += 1
            else:
                ans[pos] = nums[j] * nums[j]
                j -= 1
            pos -= 1
        return ans
                
# @lc code=end

