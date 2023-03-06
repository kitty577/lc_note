package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 中序遍历 就是递增序列
func getMinimumDifference(root *TreeNode) int {
    // 保留前一个节点的指针
    var pre *TreeNode
    min:=math.MaxInt64

    var travel func(*TreeNode)
    travel=func(node *TreeNode){
        if node==nil{
            return
        }

        travel(node.Left)

        if pre!=nil&&node.Val-pre.Val<min{
            min=node.Val-pre.Val
        }
        pre=node

        travel(node.Right)
    }

    travel(root)

    return min
}

