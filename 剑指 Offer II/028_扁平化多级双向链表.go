package main

// 类似层级遍历

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// 栈
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		var last *Node

		if node == nil {
			return nil
		}
		cur := node
		for cur != nil {
			// 记录下当前节点的下一节点
			next := cur.Next

			if cur.Child != nil {
				// 如果当前节点有孩子节点，则先插入孩子节点
				childLast := dfs(cur.Child)
				cur.Next = cur.Child
				cur.Child.Prev = cur

				if next != nil {
					childLast.Next = next
					next.Prev = childLast
				}

				cur.Child = nil
				last = childLast
			} else {
				last = cur
			}

			// 继续处理下一个节点
			cur = next
		}

		return last
	}

	dfs(root)
	return root
}
