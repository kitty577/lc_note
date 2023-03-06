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
// 递归
func trimBST(root *TreeNode, low int, high int) *TreeNode {
    if root==nil{
        return nil
    }

    // 节点值小于最小值,该节点以及其左子树都要删除===》则该节点替换成右节点
    if root.Val<low{
        right:=trimBST(root.Right,low,high)
        return right
    }

    // 节点值大于最大值,该节点以及其右子树都要删除===》则该节点替换成左节点
    if root.Val>high{
        left:=trimBST(root.Left,low,high)
        return left
    }

    root.Left=trimBST(root.Left,low,high)
    root.Right=trimBST(root.Right,low,high)
    return root
}