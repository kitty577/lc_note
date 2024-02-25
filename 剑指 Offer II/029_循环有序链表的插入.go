package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
	// 构造结点
	node := &Node{Val: x}
	// 如果原链表为空，则node构造循环即可
	if aNode == nil {
		node.Next = node
		return node
	}

	// 如果原链表只有长度为1，那么把新node加进去就行，不用考虑顺序
	if aNode.Next == aNode {
		aNode.Next = node
		node.Next = aNode
		return aNode
	}

	// 长度不为1，需要考虑有序插入了
	cur, next := aNode, aNode.Next
	for next != aNode {
		// 新node 在中间
		if x >= cur.Val && x <= next.Val {
			break
		}

		// 新node在cur 和cue.next之外的
		if cur.Val > next.Val {
			if x > cur.Val || x < next.Val {
				break
			}
		}

		cur = cur.Next
		next = next.Next
	}
	cur.Next = node
	node.Next = next
	return aNode
}
