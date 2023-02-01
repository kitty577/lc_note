package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 直径 也就是该节点的左子树的长度加上右子树的长度
// 递归
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 计算每个节点的左右孩子的高度
		l := dfs(node.Left)
		r := dfs(node.Right)

		// 路径长度为左右孩子的高度之和
		ans = max(ans, l+r)

		// 返回结点高的作为本结点的高度
		return max(l, r) + 1
	}

	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
