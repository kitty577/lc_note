package main

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	for i := 1; i < n; i++ {
		lists[i] = mergeTwo(lists[i], lists[i-1])
	}
	return lists[n-1]
}

func mergeTwo(head1, head2 *ListNode) *ListNode {
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
