#
# @lc app=leetcode.cn id=617 lang=python3
#
# [617] 合并二叉树
#
# import collections
# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution(object):
	"""广度优先搜索 + 队列

	Args:
		object ([type]): [description]
	"""
	def mergeTrees(self, t1, t2):
		"""
		:type t1: TreeNode
		:type t2: TreeNode
		:rtype: TreeNode
		"""	
	# 如果 t1和t2中，只要有一个是null，函数就直接返回
		if not (t1 and t2):
			return t2 if not t1 else t1
		queue = [(t1,t2)]
		while queue:
			r1,r2 = queue.pop(0)
			r1.val += r2.val
			# 如果r1和r2的左子树都不为空，就放到队列中
			# 如果r1的左子树为空，就把r2的左子树挂到r1的左子树上
			if r1.left and r2.left:
				queue.append((r1.left,r2.left))
			elif not r1.left:
				r1.left = r2.left
			# 对于右子树也是一样的
			if r1.right and r2.right:
				queue.append((r1.right,r2.right))
			elif not r1.right:
				r1.right = r2.right
		return t1


# @lc code=end



class Solution1:
    """深度优先搜索
    同时遍历两个二叉树， 使对应的节点合并
    
    时间复杂度： O(min(m, n)) m, n为两个二叉树的长度，最上
    空间复杂度： O(min(m, n))
    """
    def mergeTrees(self, root1: TreeNode, root2: TreeNode) -> TreeNode:
        if not root1: return root2
        if not root2: return root1
        
        merged = TreeNode(root1.val + root2.val)
        merged.left = self.mergeTrees(root1.left, root2.left)
        merged.right = self.mergeTrees(root1.right, root2.right)
        
        return merged


