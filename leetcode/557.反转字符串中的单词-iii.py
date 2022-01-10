#
# @lc app=leetcode.cn id=557 lang=python3
#
# [557] 反转字符串中的单词 III
# 双指针

# @lc code=start
class Solution:
    def reverseWords(self, s: str) -> str:
        # words = s.split()
        # res = []
        # for word in words:
        #     ws = list(word)
        #     l, r = 0, len(ws)-1
        #     while l <= r:
        #         ws[l], ws[r] = ws[r], ws[l]
        #         l += 1
        #         r -= 1
        #     ws = ''.join(ws)
        #     res.append(ws)
        # return ' '.join(res)
        
        # @lc methord2
        return " ".join(word[::-1] for word in s.split(" "))
        
# @lc code=end

