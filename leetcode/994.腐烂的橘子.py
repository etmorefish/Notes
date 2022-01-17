#
# @lc app=leetcode.cn id=994 lang=python3
#
# [994] 腐烂的橘子
#

# @lc code=start
from typing import List


class Solution:
    """BFS
    初始化队列；
    最开始的坏橘子全部入队，具体是橘子的坐标和 timetime；
    循环：当队列不为空时，先弹出队首元素，然后将这个元素能够腐烂的橘子全部入队。
    
    时间复杂度： O(n)
    空间复杂度： O(n)
    """

    def orangesRotting(self, grid: List[List[int]]) -> int:
        m, n, time = len(grid), len(grid[0]), 0
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]
        
        q = [(i, j, time) for i in range(m)
             for j in range(n) if grid[i][j] == 2]
        while q:
            i, j, time = q.pop(0)
            for di, dj in directions:
                if 0 <= i+di < m and 0 <= j+dj < n and grid[i + di][j + dj] == 1:
                    grid[i + di][j + dj] = 2
                    q.append((i + di, j + dj, time + 1))
        for row in grid:
            if 1 in row:
                return -1
        return time


# @lc code=end

class Solution1:
    """BFS
    题意：腐烂橘子到所有新鲜橘子的最短路径
    此方法下面集合不好简化代码
    """

    def orangesRotting(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])

        row = len(grid)
        col = len(grid[0])
        rotten = {(i, j) for i in range(row)
                  for j in range(col) if grid[i][j] == 2}  # 腐烂集合
        fresh = {(i, j) for i in range(row)
                 for j in range(col) if grid[i][j] == 1}  # 新鲜集合
        time = 0
        while fresh:
            if not rotten:
                return -1
            rotten = {(i + di, j + dj) for i, j in rotten for di, dj in [
                (0, 1), (0, -1), (1, 0), (-1, 0)] if (i + di, j + dj) in fresh}  # 即将腐烂的如果在新鲜的集合中，就将它腐烂
            # for i, j in rotten:
            #     for di, dj in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            #         if (i + di, j + dj) in fresh:
            #             rotten.add((i + di, j + dj))
            fresh -= rotten  # 剔除腐烂的
            time += 1
        return time


grid = [[2, 1, 1], [1, 1, 0], [0, 1, 1]]
s = Solution()
print(s.orangesRotting(grid))
