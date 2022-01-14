#
# @lc app=leetcode.cn id=567 lang=python3
#
# [567] 字符串的排列
# 滑动窗口

# @lc code=start
import collections
from typing import Counter


class Solution:
    """方法二:滑动窗口加字典
    首先统计子串每个字符串出现的次数
    创建一个与 s1 相同长度的窗口,统计窗口中元素出现的次数
    返回结果
    """
    def checkInclusion(self, s1: str, s2: str) -> bool:
        counter1 = collections.Counter(s1)
        
        n = len(s1)
        left, right = 0, n
        while right <= len(s2):
            counter2 = collections.Counter(s2[left:right])
            if counter2 == counter1:
                return True
            left += 1
            right += 1
        return False
# @lc code=end

class Solution1:
    """方法一: 字符串切片
    """
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if s1 in s2 or s1[::-1] in s2: 
            return True
        else: return False