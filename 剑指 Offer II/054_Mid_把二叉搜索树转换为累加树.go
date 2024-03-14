package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 右 中 左
// 反中序
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var reverse func(*TreeNode)
	reverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		reverse(node.Right)
		node.Val += sum
		sum = node.Val
		reverse(node.Left)
	}
	reverse(root)
	return root
}
