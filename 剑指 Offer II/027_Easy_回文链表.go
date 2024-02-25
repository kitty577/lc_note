package main

// 与上一题相似
// 可以先找到链表的中间位置，将链表后半部分进行反转操作
// 将反转后的链表与前半部分链表比较

func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
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

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	mid := findMiddle(head)
	l2 := reverse(mid.Next)
	p1, p2 := head, l2
	ans := true
	for ans && p2 != nil {
		if p1.Val != p2.Val {
			ans = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return ans
}
