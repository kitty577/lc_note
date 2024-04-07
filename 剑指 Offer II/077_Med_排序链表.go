package main

// 找到链表中点。 将问题转换为归并排序

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp, tmp1, tmp2 := dummyHead, head1, head2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val > tmp2.Val {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		} else {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		}
		tmp = tmp.Next
	}

	if tmp1 != nil {
		tmp.Next = tmp1
	}
	if tmp2 != nil {
		tmp.Next = tmp2
	}
	return dummyHead.Next
}
