package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// dfs、递归
// 二叉搜索树：如果该二叉树的左子树不为空，则左子树上所有节点的值均小于它的根节点的值；
// 若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值；它的左右子树也为二叉搜索树.

// 判断是否满足条件，需要传入当前要判断的结点、根结点的值，且大小的比较需要根据当前结点是左结点还是
// 右结点 因此递归函数设置lower、upper参数进行传参

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// 2022.10.8
