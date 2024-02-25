package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历
func findBottomLeftValue(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	curLevel := []*TreeNode{root}
	nextLevel := []*TreeNode{}
	for len(curLevel) != 0 {
		var node *TreeNode
		for _, node = range curLevel {
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
		}
		ans = node.Val
		curLevel = nextLevel
		nextLevel = nil
	}
	return ans
}
