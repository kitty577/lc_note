package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==nil{
        return nil
    }
    for{
        if root.Val>p.Val&&root.Val>q.Val{
            root=root.Left
        }
        if root.Val<p.Val&&root.Val<q.Val{
            root=root.Right
        }
        if (root.Val-p.Val)*(root.Val-q.Val)<=0{
            return root
        }
    }
    return root
}