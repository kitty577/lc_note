/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 找出偶链表、奇链表，再连接

package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	ouHead := head.Next
	jiL := head
	ouL := ouHead
	for ouL != nil && ouL.Next != nil {
		jiL.Next = ouL.Next
		jiL = jiL.Next
		ouL.Next = jiL.Next
		ouL = ouL.Next
	}
	jiL.Next = ouHead
	return head
}
