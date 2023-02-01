package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 利用好完全二叉树的性质

// 如果最后一层是满的 那么节点数为2^h-1
func countNodes(root *TreeNode) int {
    if root==nil{
        return 0
    }

    leftH,rightH:=0,0
    leftNode,rightNode:=root.Left,root.Right
    for leftNode!=nil{
        leftH++
        leftNode=leftNode.Left
    }

    for rightNode!=nil{
        rightH++
        rightNode=rightNode.Right
    }

    if leftH==rightH{  // 最底下一层是满的
        return (2<<leftH)-1
    }

    return countNodes(root.Left)+countNodes(root.Right)+1

    
}