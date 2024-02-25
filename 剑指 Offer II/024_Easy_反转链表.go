package main

// 反转
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var last *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	return last
}
