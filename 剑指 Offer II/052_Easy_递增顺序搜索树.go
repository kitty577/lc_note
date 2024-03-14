package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树特点： 左子树 》 根结点 》 右子树
func increasingBST(root *TreeNode) *TreeNode {
	dummyNode := &TreeNode{}
	resNode := dummyNode

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left) // 左

		// 中 处理节点
		resNode.Right = node
		node.Left = nil
		resNode = node

		inorder(node.Right) // 右
	}

	inorder(root)
	return dummyNode.Right
}
