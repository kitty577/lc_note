package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

// 深度优先搜索
func dfs(root *TreeNode, preSun int) int {
	if root == nil {
		return 0
	}
	sum := preSun*10 + root.Val
	if root.Left == nil && root.Right == nil { // 走到末尾叶子节点了
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}
