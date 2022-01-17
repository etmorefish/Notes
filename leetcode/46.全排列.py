#
# @lc app=leetcode.cn id=46 lang=python3
#
# [46] 全排列
#

# @lc code=start
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        def dfs(nums, tmp):
            if not nums: 
                res.append(tmp)
                return
            for i in range(len(nums)):
                dfs(nums[:i] + nums[i+1:], tmp +[nums[i]])
        dfs(nums, [])
        return res
# @lc code=end

