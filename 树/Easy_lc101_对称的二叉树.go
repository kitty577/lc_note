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

func isSymmetric(root *TreeNode) bool {
    if root==nil{
        return true
    }else if root.Left==nil&&root.Right==nil{
        return true
    }else if root.Left==nil||root.Right==nil{
        return false
    }else {
        return check(root.Left,root.Right)
    } 
}

func check(p,q *TreeNode)bool{
    if p==nil&&q==nil{
        return true
    }else if p==nil||q==nil{
        return false
    }

    return p.Val==q.Val&&check(p.Left,q.Right)&&check(p.Right,q.Left)
}

// 2023.01.02