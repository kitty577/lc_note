package main

// 反转链表

// 1-->2-->3-->4-->5

func reverseList(head *ListNode) *ListNode {
	var last *ListNode
	for head != nil {
		headNext := head.Next
		head.Next = last
		last = head
		head = headNext
	}
	return last
}
