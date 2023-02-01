// 思路：
// 方法一：可以借助map,标记已经访问过的结点。第一次走到访问过的结点，则为入环点
// 方法二：快慢指针：slow走1步 fast走2步
//        起点到入环点距离为a,入环点到相遇点距离为b，相遇点回到入环点距离为c
//        fast：a+n(b+c)+b
//        slow: a+b
//        由与fast=2*slow ==>a=c+(n-1)(b+c)  ==>a==c 则只需要在相遇时再定义一个指针 与slow指针同步走，当两者相遇的时候的结点 就是入环点

package main

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			pre := head
			for pre != slow {
				pre = pre.Next
				slow = slow.Next
			}
			return pre
		}
	}
	return nil
}

// 2022.10.3
