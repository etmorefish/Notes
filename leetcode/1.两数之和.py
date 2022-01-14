#
# @lc app=leetcode.cn id=1 lang=python3
#
# [1] 两数之和
#

# @lc code=start
from typing import List


class Solution:
    """方法一: hashmap
    将列表元素存到hash表里, 
    遍历列表,用 target 减去当前的值,得到两数之和的另一个数,
    判断这个数是否在hash表里,和是否和当前元素相等
    
    """
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashmap = {}
        for i, v in enumerate(nums):
            hashmap[v] = i
        for i, v in enumerate(nums):
            another =  hashmap.get(target - v)
            if another is not None and another != i:
                return [i, another]
            
# @lc code=end

nums = [3, 2, 4]
target = 6
s = Solution()
print(s.twoSum(nums, target))