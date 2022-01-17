#
# @lc app=leetcode.cn id=542 lang=python3
#
# [542] 01 矩阵
#

# @lc code=start
from typing import List


class Solution:
    """广度优先搜索
    对于矩阵中的每一个元素，如果它的值为 0，
    那么离它最近的 0 就是它自己。
    如果它的值为 1，那么我们就需要找出离它最近的 0，
    并且返回这个距离值。那么我们如何对于矩阵中的每一个 1，
    """
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        m, n = len(mat), len(mat[0])
        dist = [[0]*n for _ in range(m)]
        
        zero_pos = [(i, j) for i in range(m) for j in range(n) if mat[i][j] == 0]
        q = zero_pos
        seen = set(zero_pos)
        
        while q:
            i, j = q.pop(0)
            for ni, nj in [(i - 1, j), (i + 1, j), (i, j - 1), (i, j + 1)]:
                if 0 <= ni < m and 0 <= nj < n and (ni, nj) not in seen:
                    dist[ni][nj] = dist[i][j] + 1
                    q.append((ni, nj))
                    seen.add((ni, nj))
        
        return dist


# @lc code=end

mat = [[0,0,0],[0,1,0],[0,0,0]]
s = Solution()
s.updateMatrix(mat)