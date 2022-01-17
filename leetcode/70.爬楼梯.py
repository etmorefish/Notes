#
# @lc app=leetcode.cn id=70 lang=python3
#
# [70] 爬楼梯
#

# @lc code=start
class Solution:
    def climbStairs(self, n: int) -> int:
        if n == 0 or n == 1:
            return 1
        elif n == 2:
            return 2
        else:
            ret = 0
            a = 1
            b = 2
            for i in range(2, n):
                ret = a + b
                a = b
                b = ret
            return ret
        
# @lc code=end

