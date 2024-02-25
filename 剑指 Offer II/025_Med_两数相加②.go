package main

// 末位开始相加 并且有进位。
// 想到从末尾开始遍历。---》 反转链表
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	rl1 := reverse(l1)
	rl2 := reverse(l2)
	carry := 0 // 进位
	var ans *ListNode
	for rl1 != nil || rl2 != nil || carry != 0 {
		sum := 0
		if rl1 != nil {
			sum += rl1.Val
			rl1 = rl1.Next
		}
		if rl2 != nil {
			sum += rl2.Val
			rl2 = rl2.Next
		}
		sum += carry
		carry = sum / 10
		sum = sum % 10
		node := &ListNode{Val: sum, Next: ans}
		ans = node
	}
	return ans
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
