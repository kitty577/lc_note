package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜素树中序遍历得到的就是升序数组
// 使用双指针 两数之和
func findTarget(root *TreeNode, k int) bool {
	nums := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			return true
		} else if sum > k {
			right--
		} else if sum < k {
			left++
		}
	}
	return false
}
