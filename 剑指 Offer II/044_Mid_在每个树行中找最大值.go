package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历
func largestValues(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	curLevel := []*TreeNode{root}
	nextLevel := []*TreeNode{}

	for len(curLevel) > 0 {
		value := curLevel[0].Val

		for _, n := range curLevel {
			value = max(value, n.Val)
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
		}

		ans = append(ans, value)
		curLevel = nextLevel
		nextLevel = nil
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
