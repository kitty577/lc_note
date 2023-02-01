package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// 前序遍历
func preorder(root *Node) (ans []int) {
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, ch := range node.Children {
			dfs(ch)
		}
	}
	dfs(root)
	return
}

// 2022.10.8
