package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢双指针

// 倒数第N个 双指针，只需要让快指针先走n步，当快指针走到末尾，则慢指针就在倒数第N个的位置上
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	slow, fast := dummyHead, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
