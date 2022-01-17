#
# @lc app=leetcode.cn id=77 lang=python3
#
# [77] 组合
#

# @lc code=start



from typing import List


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:

        res=[]  #存放符合条件结果的集合
        tmp=[] #
        
        def dfs(start, level, tmp): 
            if n-start+1 < level:
                return 
            if level == 0:
                res.append(tmp[::])
            
            for i in range(start, n+1):
                tmp.append(i)
                dfs(i+1, level -1, tmp)
                tmp.pop()
        dfs(1, k, tmp)
        return res
    
        
# @lc code=end

n = 4
k = 2
s = Solution()
s.combine(n, k)