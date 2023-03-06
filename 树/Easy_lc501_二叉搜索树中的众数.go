package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉搜索树中序遍历
func findMode(root *TreeNode) []int {
    res:=make([]int,0)
    count:=1
    max:=1
    var pre *TreeNode
    var travel func(*TreeNode)
    travel=func(node *TreeNode){
        if node==nil{
            return 
        }

        travel(node.Left)

        if pre!=nil&&pre.Val==node.Val{
            count++
        }else{
            count=1
        }

        if count>=max{
            if count==max{
                res=append(res,node.Val)
            }else{
                res=[]int{node.Val}
            }

            max=count
        }
        pre=node

        travel(node.Right)
    }

    travel(root)
    return res
}