#
# @lc app=leetcode.cn id=876 lang=python3
#
# [876] 链表的中间结点
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    """方法三 双指针 快慢指针
    定义两个指针变量 slow， fast， 初始位置都位于第一个节点，
    slow 指针每次走一步，fast 指针每次走两步，同时走
    当 fast 指针走到链表末尾时，慢指针 slow 就走到了链表的中间位置

    时间复杂度：O(N)O(N)，其中 NN 是给定链表的结点数目。

    空间复杂度：O(1)O(1)，只需要常数空间存放 slow 和 fast 两个指针。
    Returns:
        [type]: [description]
    """
    def middleNode(self, head: ListNode) -> ListNode:
        slow, fast = head, head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            
        return slow

# @lc code=end


class Solution3:
    """方法二：单指针
    第一次遍历，得到链表的长度n，
    第二次遍历一半，即得到链表中间的值
    
    时间复杂度：O(N)，其中 N 是给定链表中的结点数目。

    空间复杂度：O(1)，只需要常数变量的指针。
    """
    def middleNode(self, head: ListNode) -> ListNode:
        n, cur = 0, head
        while cur:
            n += 1
            cur = cur.next
        m, cur = 0, head
        while m < n//2:
            m += 1
            cur = cur.next
        return cur
        
class Solution2:
    """方法一：数组
    
    此方法运行超时
    
    对链表中的元素进行遍历，然后一次存入列表中
    
    时间复杂度：O(N)O(N)，其中 NN 是给定链表中的结点数目。

    空间复杂度：O(N)O(N)，即数组 A 用去的空间。
    """
    def middleNode(self, head: ListNode) -> ListNode:
        nums = [head]
        while nums[-1].next:
            nums.append(head.next)
        return nums[len(nums) // 2]

