package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root != nil {
		if root.Val < key { // 根据二叉搜索树特点，缩小搜索范围
			root.Right = deleteNode(root.Right, key)
		} else if root.Val > key { // 根据二叉搜索树特点，缩小搜索范围
			root.Left = deleteNode(root.Left, key)
		} else if root.Val == key { // 找到与key值相同的节点
			if root.Left == nil || root.Right == nil {
				// 左不空右空 左子树补位
				if root.Left != nil {
					root = root.Left
				} else {
					// 左空右不空 右子树补位
					root = root.Right
				}
			} else { // 左右均不为空,删除节点的左子树放在删除节点的右子树的最左边节点上
				node := root.Left
				for node.Right != nil {
					node = node.Right
				}
				node.Left = deleteNode(root.Left, node.Val)
				node.Right = root.Right
				root = node
			}
		}
	}
	return root
}
