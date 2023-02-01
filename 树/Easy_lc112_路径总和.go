package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 方法一：回溯
func hasPathSum1(root *TreeNode, targetSum int) bool {
    if root==nil{
        return false
    }

    return backtracking(root,targetSum)
}

func backtracking(node *TreeNode,remain int)bool{
    // 走到了叶子节点 并满足和的要求
    if node.Left==nil&&node.Right==nil&&remain==0{ 
        return true
    }

    // 走到了叶子节点 但不满足和
    if node.Left==nil&&node.Right==nil{
        return false
    }

    // 未走到叶子节点

    // 开始处理左子树
    if node.Left!=nil{
        remain=remain-node.Left.Val
        if backtracking(node.Left,remain){
            return true
        }
        // 还原现场
        remain=remain+node.Left.Val
    }

    // 开始处理右子树
    if node.Right!=nil{
        remain=remain-node.Right.Val
        if backtracking(node.Right,remain){
            return true
        }
        // 还原现场
        remain=remain+node.Right.Val
    }

    return false
}


// 方法二：递归
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root==nil{
        return false
    }
    targetSum=targetSum-root.Val
    if root.Left==nil&&root.Right==nil&&targetSum==0{
        return true
    }

    return hasPathSum(root.Left,targetSum)||hasPathSum(root.Right,targetSum)
}