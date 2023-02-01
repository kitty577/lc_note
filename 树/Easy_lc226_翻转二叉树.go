package main

// 递归
// 从根节点开始，递归的对树进行遍历，并从叶子节点开始翻转
// 如果当前遍历的结点root的左右两棵子树都已经翻转。那么我们只要交换两棵子树的位置，
// 即可完成以root为根结点的整棵子树的翻转

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Right)
	right := invertTree(root.Left)
	root.Left = left
	root.Right = right
	return root
}

// 2022.10.26
