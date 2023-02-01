// 思路：层序遍历 bfs
// 借助队列，记录下每一层的结点
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	// res 存储结果
	res := [][]int{}
	if root == nil {
		return res
	}
	// currentLevel存储当前处理的层的结点
	currentLevel := []*TreeNode{root}
	// 处理当前层
	for len(currentLevel) > 0 {
		// 存储当前层的值
		levelVal := []int{}
		// 根据当前层，收集下一层的节点
		nextLevel := []*TreeNode{}

		for _, node := range currentLevel {

			levelVal = append(levelVal, node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		// 将当前层的数值存放到二维数组中
		res = append(res, levelVal)

		//将下一层变成当前层 继续处理
		currentLevel = nextLevel
	}
	return res
}

// 2022.10.8
