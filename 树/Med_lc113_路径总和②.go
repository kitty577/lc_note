package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}
	path := []int{root.Val}
	backtracking(root, targetSum-root.Val, &path, &ans)
	return ans
}

func backtracking(node *TreeNode, count int, path *[]int, res *[][]int) {
	// 走到叶子节点 并且值满足要求，更新ans
	if node.Left == nil && node.Right == nil && count == 0 {
		*res = append(*res, append([]int{}, *path...))
		return
	}

	// 遍历完 值还是不符合要求
	if node.Left == nil && node.Right == nil {
		return
	}

	// 开始处理左
	if node.Left != nil {
		*path = append(*path, node.Left.Val)
		count = count - node.Left.Val
		backtracking(node.Left, count, path, res)
		// 还原现场
		*path = (*path)[:len(*path)-1]
		count = count + node.Left.Val
	}

	// 开始处理右
	if node.Right != nil {
		*path = append(*path, node.Right.Val)
		count = count - node.Right.Val
		backtracking(node.Right, count, path, res)
		// 还原现场
		*path = (*path)[:len(*path)-1]
		count = count + node.Right.Val
	}
}
