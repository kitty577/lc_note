package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 找到根节点相同的 再比较树是否相同
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root==nil{
        return false
    }
    return check(root,subRoot)||isSubtree(root.Left,subRoot)||isSubtree(root.Right,subRoot)
}

func check(p,q *TreeNode)bool{
    if p==nil&&q==nil{
        return true
    }else if p==nil||q==nil{
        return false
    }else{
        return p.Val==q.Val&&check(p.Left,q.Left)&&check(p.Right,q.Right)
    }
}