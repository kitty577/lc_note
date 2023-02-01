package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 2022.12.11

// 二叉树的最大/最小深度
// 思路一：(自底向上统计信息，分治思想)
// 		最大深度=max(左子树最大深度，右子树最大深度)+1
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth1(root.Left)
	right := maxDepth1(root.Right)

	return max(left, right) + 1
}

// 思路二：（自顶向下维护信息）
// 		把深度作为一个全局变量-----一个跟随结点移动而动态改变的信息
// 		递归一层，变量+1，在叶子处更新结点
// 		！！！这种写法要注意还原现场
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dep := 1 // 根结点 深度1
	ans := 0 // 最终答案

	find(root, &dep, &ans)
	return ans
}

func find(root *TreeNode, dep *int, ans *int) {
	if root == nil {
		return
	}
	*ans = max(*ans, *dep) // 更新答案
	*dep++                 // 开始遍历下一层 则深度+1
	find(root.Left, dep, ans)
	find(root.Right, dep, ans)
	*dep-- // 还原现场
}

func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dep := 1 // 根结点 深度1
	ans := 0 // 最终答案
	var find func(*TreeNode)
	find = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = max(ans, dep) // 更新答案
		dep++               // 开始遍历下一层 则深度+1
		find(root.Left)
		find(root.Right)
		dep-- // 还原现场
	}
	find(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
