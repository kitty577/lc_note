package main

// 有序数组nums是按升序排列。因此只要找到nums中点，以中点为根结点，左边为左子树，右边为右子树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left < right {
		return nil
	}
	mid := (left + right) / 2
	root := *&TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return &root
}

// 2022.12.22
