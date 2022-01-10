#
# @lc app=leetcode.cn id=19 lang=python3
#
# [19] 删除链表的倒数第 N 个结点
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    """方法三：快慢指针
    定义slow 和 fast 两个快慢指针，fast 先走 n 步， 然后 slow 指针和 fast 同时走，
    直到 fast 指针走到末尾，此时 slow位于倒数第 n 个节点
    
    时间复杂度：O(L)O(L)，其中 LL 是链表的长度。

    空间复杂度：O(1)O(1)。
    """ 
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        dummy = ListNode(0, head)
        slow, fast =dummy, head
        
        for i in range(n):
            fast = fast.next
        
        while fast:
            fast = fast.next
            slow = slow.next
        slow.next = slow.next.next
        return dummy.next
        
        
        
# @lc code=end

class Solution:
    """方法二：栈
    我们也可以在遍历链表的同时将所有节点依次入栈。根据栈「先进后出」的原则，
    我们弹出栈的第 nn 个节点就是需要删除的节点，并且目前栈顶的节点就是待删除节点的前驱节点。这样一来，删除操作就变得十分方便了。
    
    时间复杂度：O(L)O(L)，其中 LL 是链表的长度。

    空间复杂度：O(L)O(L)，其中 LL 是链表的长度。主要为栈的开销。
    """ 
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        dummy = ListNode(0, head)
        stack = list()
        cur = dummy
        while cur:
            stack.append(cur)
            cur = cur.next
            
        for i in range(n):
            stack.pop()
        
        prev = stack[-1]
        prev.next = prev.next.next
        
        return dummy.next
            
            


class Solution1:
    """方法一：计算链表长度
    一种容易想到的方法是，我们首先从头节点开始对链表进行一次遍历，得到链表的长度 LL。随后我们再从头节点开始对链表进行一次遍历，当遍历到第 L-n+1L−n+1 个节点时，它就是我们需要删除的节点。

    为了与题目中的 nn 保持一致，节点的编号从 1 开始，头节点为编号 1 的节点。

    为了方便删除操作，我们可以从哑节点开始遍历 L-n+1 个节点。当遍历到第 L-n+1 个节点时，它的下一个节点就是我们需要删除的节点，这样我们只需要修改一次指针，就能完成删除操作。

    时间复杂度：O(L)O(L)，其中 LL 是链表的长度。

    空间复杂度：O(1)O(1)。
    """
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        def getLength(head: ListNode):
            length = 0
            while head:
                length += 1
                head = head.next
            return length
        dummy = ListNode(0, head)
        length = getLength(head)
        cur = dummy
        for i in range(1, length - n + 1):
            cur = cur.next
        cur.next = cur.next.next
        return dummy.next

