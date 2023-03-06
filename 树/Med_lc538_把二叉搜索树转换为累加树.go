package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 二叉搜索树 中序遍历就是从小到大的递增序列
// 反中序
var pre int
func convertBST(root *TreeNode) *TreeNode {
    pre=0
    traversal(root)
    return root
}

func traversal(cur *TreeNode){
    if cur==nil{
        return
    }
    traversal(cur.Right)
    cur.Val+=pre
    pre=cur.Val
    traversal(cur.Left)
}