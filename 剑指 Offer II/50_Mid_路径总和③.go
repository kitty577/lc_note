package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs+ 前缀和
func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	preSum := make(map[int64]int)
	preSum[0] = 1

	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, cur int64) {
		if node == nil {
			return
		}
		cur += int64(node.Val)
		res += preSum[cur-int64(targetSum)]
		preSum[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)

		preSum[cur]--
	}

	dfs(root, 0)

	return res
}
