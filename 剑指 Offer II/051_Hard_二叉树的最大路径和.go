package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 路径和= 左子树贡献的结点和 + 右子树贡献的结点和 + 当前结点值
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := root.Val
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		// 更新一次答案，可能本次路径就是最大
		sum := leftGain + rightGain + node.Val
		maxSum = max(maxSum, sum)

		// 返回结点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
