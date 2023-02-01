package main

// 递归
// 删除所有节点值为0的子树，可以从叶子节点开始删除，自底向上递归
// ====》 题目转化为 删除所有值为0的叶子节点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}

// 2022.12.13
