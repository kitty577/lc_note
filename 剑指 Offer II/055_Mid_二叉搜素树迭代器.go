package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nums []int
}

func Constructor(root *TreeNode) BSTIterator {
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
	return BSTIterator{nums: nums}
}

func (this *BSTIterator) Next() int {
	top := this.nums[0]
	this.nums = this.nums[1:]
	return top
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nums) > 0
}
