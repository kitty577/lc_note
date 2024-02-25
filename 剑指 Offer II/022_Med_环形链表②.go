package main

// 快慢双指针
// 慢指针走一步，快指针走两步。如果有环，则一定会在环内相遇
// 设head到入环点之间的举例为a，入环点到环内相遇点的距离为b,环内相遇点回到入环点的距离为c
// 慢指针走的距离为 a+b
// 快指针走的距离为 a+n(b+c)+b
// 2*slow=fast
// ==> 2a+2b=a+nb+nc+b
// ==> a=nb+nc-b
// ==> a=(n-1)(b+c)+c
// ==> a=c
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		// 快慢指针往前走
		slow = slow.Next
		fast = fast.Next.Next
		// 当快慢指针相遇，根据推测出的，a=c----头到入环点的距离等于相遇点回到入环点的距离
		if fast == slow {
			meet := head
			for meet != slow {
				meet = meet.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
