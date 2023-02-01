package main

// 利用二叉搜索树特点 中序排序后就是从小到大的顺序

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 二叉搜索树 左>根>右
// 使用中序遍历，则遍历结果就是从小到大的排序
func kthSmallest(root *TreeNode, k int) int {
	nums := []int{}
	var inner func(*TreeNode)
	inner = func(node *TreeNode) {
		if node == nil {
			return
		}
		inner(node.Left)
		nums = append(nums, node.Val)
		inner(node.Right)
	}
	inner(root)

	return nums[k-1]
}

func kthSmallest2(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
