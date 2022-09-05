package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
