package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉搜索树特点:左《根《右
func searchBST(root *TreeNode, val int) *TreeNode {
    if root==nil{
        return nil
    }

    if root.Val==val{
        return root
    }

    if root.Val<val{
        return searchBST(root.Right,val)
    }

    if root.Val>val{
        return searchBST(root.Left,val)
    }
    return nil
}

// 2023.02.01