package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 迭代
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	stack := []*TreeNode{root}
	paths := []string{""}
	res := []string{}

	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		path := paths[l-1]
		stack = stack[:l-1]
		paths = paths[:l-1]

		if node.Left == nil && node.Right == nil { // 已经处理到了叶子节点
			res = append(res, path+strconv.Itoa(node.Val))
			//continue
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}
	return res
}

// 递归
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}

	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			res = append(res, v)
			return
		}

		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}

	travel(root, "")
	return res
}
