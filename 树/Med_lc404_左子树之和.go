package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归写法
func sumOfLeftLeaves(root *TreeNode) int {
    if root==nil{
        return 0
    }
    leftValue:=sumOfLeftLeaves(root.Left)

    if root.Left!=nil&&root.Left.Left==nil&&root.Left.Right==nil{
        leftValue=root.Left.Val
    }
    rightValue:=sumOfLeftLeaves(root.Right)

    return leftValue+rightValue
}

// 迭代
func sumOfLeftLeaves2(root *TreeNode) int {
    if root==nil{
        return 0
    }
    st:=[]*TreeNode{root}
    result:=0

    for len(st)>0{
        node:=st[len(st)-1]
        st=st[:len(st)-1]
        if node.Left!=nil&&node.Left.Left==nil&&node.Left.Right==nil{
            result+=node.Left.Val
        }

        if node.Left!=nil{
            st=append(st,node.Left)
        }

        if node.Right!=nil{
            st=append(st,node.Right)
        }
    }

    return result
}