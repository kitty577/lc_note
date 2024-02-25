package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	curLevel := []*TreeNode{root}
	nextLevel := []*TreeNode{}
	for len(curLevel) > 0 {
		var node *TreeNode
		for _, node = range curLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		ans = append(ans, node.Val)
		curLevel = nextLevel
		nextLevel = nil
	}
	return ans
}
