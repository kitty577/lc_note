package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathFrom(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)

}

func pathFrom(node *TreeNode, count int) int {
	if node == nil {
		return 0
	}

	res := 0
	// 如果该节点的值刚好等于与目标值还差的值
	if node.Val == count {
		res++
	}

	// 这里要注意 由于题中说的路径无须从根节点开始 也无须在叶子节点结束，所以不管是否已经找到满意的路径，都要继续走下去
	// 可能存在正负值抵消的情况
	res += pathFrom(node.Left, count-node.Val)
	res += pathFrom(node.Right, count-node.Val)
	return res
}

// 2022.12.22
