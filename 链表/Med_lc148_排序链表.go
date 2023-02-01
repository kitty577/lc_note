/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 归并排序
// 找到链表的中点，以中点为分界，将链表拆分成两个子链表。
// 寻找链表的中点可以使用快慢指针的做法，快指针每次移动2 步，慢指针每次移动1 步，当快指针到达链表末尾时，慢指针指向的链表节点即为链表的中点。

// 对两个子链表分别排序。

// 将两个排序后的子链表合并，得到完整的排序后的链表。

// 链表截断！！注意head tail（为nil的情况）

package main

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
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

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

// 2022.10.21
