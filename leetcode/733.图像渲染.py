#
# @lc app=leetcode.cn id=733 lang=python3
#
# [733] 图像渲染
#


# @lc code=start
from typing import List
from queue import Queue


class Solution:
    """DFS 使用栈
    
    """
    def floodFill(self, image: List[List[int]], sr: int, sc: int, newColor: int) -> List[List[int]]:
        if newColor == image[sr][sc]: return image
        stack, old = [(sc, sr)], image[sr][sc]
        while stack:
            point = stack.pop()
            image[point[0]][point[1]] = newColor
            for new_i, new_j in zip((point[0], point[0], point[0] + 1, point[1] - 1), (point[1] + 1, point[1] - 1, point[1], point[1])):
                if 0 <= new_i < len(image) and 0 <= new_j < len(image[0]) and image[new_i][new_j] == old:
                    stack.append((new_i, new_j))
        return image
        


# @lc code=end

class Solution2:
    """优化 Solution1  遍历改成 zip
    
    """
    def floodFill(self, image: List[List[int]], sr: int, sc: int, newColor: int) -> List[List[int]]:
        if newColor == image[sr][sc]: return image
        que, old = [(sr, sc)], image[sr][sc]
        while que:
            point = que.pop()
            image[point[0]][point[1]] = newColor
            for new_i, new_j in zip((point[0], point[0], point[0] - 1, point[1] + 1), (point[1] + 1, point[1] - 1, point[1], point[1])):
                if 0 <= new_i < len(image) and 0 <= new_j < len(image[0]) and image[new_i][new_j] == old:
                    que.insert(0, (new_i, new_j))
        return image
                

image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1
sc = 1 
newColor = 2
s = Solution()
s.floodFill(image, sr, sc, newColor)

class Solution1:
    """
    BFS
    BFS，即 广度优先搜索。

    思路
    首先找到初始节点，给它染色，这个初始节点当作第一层。
    找到初始节点周围四个节点，给它们染色（符合条件的才能染），这四个节点当作第二层。
    再找到这四个节点周围八个节点，给它们染色，这八个节点当作第三层。
    重复以往，层层递进，直到找不到符合要求的节点。
    思路很好理解对吧，就是一个从中间向外扩散的过程。可是怎么实现呢？现在给您键盘，恐怕还写不出。

    实现
    这里就要介绍一下 队列，因为 广度优先搜索 和 队列 是好基友。
    什么是队列？就是一个先进先出的数组，和我们日常生活中的排队很像。
    当我们向队列插入一个新数的时候，它插在最后，当我们取出一个数的时候，要从头取。

    补充——关于在 Python 中使用队列
    在 Python 中，可以使用以下几种方法实现队列

    collections包里的deque，对应操作
    pop()从尾取出
    appendleft() 从头插入
    queue包中的queue，对应操作
    put() 插入
    get() 取出
    直接使用list，只要保证只使用
    pop() 取出
    insert(0,) 插入
    或者只使用
    append() 插入
    list[0]并且del list[0] 取出
    两者使用list方法的不同就区别于你把哪个当头，哪个当尾
    三种方法各有优劣。

    第一种是正统的Python的双端队列，缺点是调用的函数有点复杂，可能一不小心写了append，就不对了。
    第二种使用封装的函数很直接，put()和get()不容易搞混淆。但是queue类型其实里面本身就装了一个deque，有点脱裤子放X的感觉。
    第三种优势在于不用调包，但是函数使用逻辑可能造成混淆。在
    这里，完整版代码采用第二种，好理解，精简版代码采用第三种，省行数。三种方式可以按照你的喜好互相替换，完全不影响结果。

    我们可以这样利用 队列 实现 广度优先搜索。

    我们设置一个队列，先把初始点添加进去
    规定每次从队列取出一个坐标
    对这个坐标染色，并且把这个坐标的邻居（符合要求且不重复的好邻居），放到队列中。
    当这个队列为空的时候，说明染色完成
    因为队列每次取出的是最后的，而每次添加的是放在最前面，所以可以想象到，每次先处理的都是层级最少的，最接近初始点的，然后慢慢扩大，这样就实现了 广度优先搜索。

    在这道题目里，层级是 不重要 的，这也是为什么后来还有个深度优先搜索，也可以解决这道题目。但是广度优先搜索的特点就在于这个层级，在很多题目当中它是很重要的。
    有时候，对队列取出元素的时候，要设置循环，固定住当前的队列项，对当前的队列项操作——因为当前的队列项一定是相同层级的。
    例如，在我们寻找到达某个节点的最小步数的时候，层级也就是步数就显得尤为重要了。
    在 leetcode 当中，有很多题都是需要 广度优先搜索 的，这是一种解题的思想，要熟练掌握。而实现这个思想，又离不开 队列。两者相辅相成。
    在这道题目当中，要注意起始颜色和目标颜色一定要不同，不然会死循环！

    """
    def floodFill(self, image: List[List[int]], sr: int, sc: int, newColor: int) -> List[List[int]]:
        # 起始颜色和目标颜色相同,直接返回元图
        if newColor == image[sr][sc]: return image
        
        # 设置四个方向的偏移量, 省事技巧 上下左右
        directions = {(1, 0), (-1, 0), (0, -1), (0, 1)}
        # 记录起始颜色
        originalColor = image[sr][sc]
        
        # 构造一个队列,把起始点放进去
        que = Queue()
        que.put((sr, sc))
        # 当队列为空跳出循环
        while not que.empty():
            # 取出队列中的点,并染色(修改)
            point = que.get()
            image[point[0]][point[1]] = newColor
            # 遍历四个方向
            for direction in directions:
                # 产生新点:(new_i, new_j)
                new_i = point[0] + direction[0]
                new_j = point[1] + direction[1]
                # 定义边界,如果这个点在区域内,并且他和原来的颜色相同
                if 0 <= new_i < len(image) and 0 <= new_j < len(image[0]) and image[new_i][new_j] == originalColor:
                    que.put((new_i, new_j))
        return image
            
        
        

