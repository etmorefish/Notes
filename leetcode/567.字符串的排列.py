#
# @lc app=leetcode.cn id=567 lang=python3
#
# [567] 字符串的排列
# 滑动窗口

# @lc code=start
class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if s1 in s2 or s1[::-1] in s2: 
            return True
        else: return False
# @lc code=end

