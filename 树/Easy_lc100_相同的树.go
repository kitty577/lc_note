package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p==nil&&q==nil{
        return true
    }else if p==nil||q==nil{
        return false
    }else if p.Val!=q.Val{
        return false
    }else{
        return isSameTree(p.Left,q.Left)&&isSameTree(p.Right,q.Right)
    }
}

// 2023.01.02