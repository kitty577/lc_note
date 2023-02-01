package main

// 快慢指针 当快指针走N步后慢指针开始走====》当快指针走到尾部时，慢指针刚好走到要删除的点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
