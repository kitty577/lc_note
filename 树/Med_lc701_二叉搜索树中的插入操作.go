package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉搜索树特点
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root==nil{  // 到了叶子节点了
        node:=&TreeNode{val,nil,nil}
        return node
    }
    if root.Val>val{
        root.Left=insertIntoBST(root.Left,val)
    }
    if root.Val<val{
        root.Right=insertIntoBST(root.Right,val)
    }
    return root
}