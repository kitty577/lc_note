package main

// 1.各自遍历，得到两链表的长度。长链表先走差值
//  2. 两指针同时分别从两链表头部开始走。由于相交后面的路程相同，所以只需要走对方相遇前走过的路。
//     “我走到终点也没见到你，于是走过你来时的路，等到我们相遇时才发现，你也走过我来时的路””
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else if curB == nil {
			curB = headA
		} else {
			curA = curA.Next
			curB = curB.Next
		}
	}
	return curA
}
