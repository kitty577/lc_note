package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 公共祖先 想到的是自底向上查==》回溯
// 对树而言 左右中 即是回溯过程
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root==nil{
        return nil
    }

    if root==p||root==q{
        return root
    }

    left:=lowestCommonAncestor(root.Left,p,q)
    right:=lowestCommonAncestor(root.Right,p,q)

    if left!=nil&&right!=nil{
        return root
    }
    if left==nil{
        return right
    }
    if right==nil{
        return left
    }

    return nil
}