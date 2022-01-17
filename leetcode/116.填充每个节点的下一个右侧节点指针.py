#
# @lc app=leetcode.cn id=116 lang=python3
#
# [116] 填充每个节点的下一个右侧节点指针
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""


class Solution:
    """
    :type root: Node
    :rtype: Node
    """

    def connect(self, root: 'Optional[Node]') -> 'Optional[Node]':
        def dfs(root):
            if not root:
                return
            left = root.left
            right = root.right
            # 配合动画演示理解这段，以root为起点，将整个纵深这段串联起来
            while left:
                left.next = right
                left = left.right
                right = right.left
            # 递归的调用左右节点，完成同样的纵深串联
            dfs(root.left)
            dfs(root.right)
        dfs(root)
        return root


# @lc code=end


class Solution2:
    """error
    """

    def connect(self, root: 'Optional[Node]') -> 'Optional[Node]':
        if root is None:
            return root

        pre = root
        # 循环条件是当前节点的left不为空，当只有根节点
        # 或所有叶子节点都出串联完后循环就退出了

        while pre.left:
            tmp = pre
            while tmp:
                # 将tmp的左右节点都串联起来
                # 注:外层循环已经判断了当前节点的left不为空
                tmp.left.next = tmp.right
                # 下一个不为空说明上一层已经帮我们完成串联了
                if tmp.next:
                    tmp.right.next = tmp.next.left
                    # 继续右边遍历
                tmp = tmp.next
                # 从下一层的最左边开始遍历
            pre = pre.left
        return root


class Solution1:
    """BFS + 栈
    二叉树的层次遍历，每次临时遍历的节点都会保存到一个队列中
    时间复杂度： O(n)
    空间复杂度： O(n)
    """

    def connect(self, root: 'Optional[Node]') -> 'Optional[Node]':
        if not root:
            return root
        queue = [root]

        while queue:
            size = len(queue)
            tmp = queue[0]
            for i in range(1, size):
                tmp.next = queue[i]
                tmp = queue[i]
            for i in range(size):
                tmp = queue.pop(0)
                if tmp.left:
                    queue.append(tmp.left)
                if tmp.right:
                    queue.append(tmp.right)

        return root
