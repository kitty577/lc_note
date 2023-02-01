package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 自底向上
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minD, minDepth1(root.Left))
	}
	if root.Right != nil {
		minD = min(minD, minDepth1(root.Right))
	}

	return minD + 1
}

// 自顶向下
 func minDepth(root *TreeNode) int {
    if root==nil{
        return 0
    }

    currentLevelNode:=[]*TreeNode{root}
    ans:=1

    for len(currentLevelNode)!=0{
        nextLevelNode:=[]*TreeNode{}
        for _,node:=range currentLevelNode{
            if node.Left==nil&&node.Right==nil{
                return ans
            }
            if node.Left!=nil{
                nextLevelNode=append(nextLevelNode,node.Left)
            }
            if node.Right!=nil{
                nextLevelNode=append(nextLevelNode,node.Right)
            }
        }

        currentLevelNode=nextLevelNode
        ans++
    }
    return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 2022.12.11
