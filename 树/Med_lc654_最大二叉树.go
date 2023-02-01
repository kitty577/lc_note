package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归构建

// 找到数组中的最大值和对应的下标，以此最大值构造根节点，划分数组
// 最大值所在下标以左的区间 构造左子树
// 最大值所在下标以右的区间 构造左子树
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums)==0{
        return nil
    }
    // 最大值构建根节点
    index:=findMaxIdx(nums)
    root:=&TreeNode{nums[index],nil,nil}
    root.Left=constructMaximumBinaryTree(nums[:index])
    root.Right=constructMaximumBinaryTree(nums[index+1:])

    return root
}


func findMaxIdx(nums []int)int{
    index:=0
    for i,v:=range nums{
        if nums[index]<v{
            index=i
        }
    }
    return index
}


// 2023.02.01