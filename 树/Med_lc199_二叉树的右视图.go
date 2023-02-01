package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 可以层序遍历，每一层的最后的那个数就是这一层右视看到的数
// 层序遍历，借助队列
func rightSideView(root *TreeNode) []int {
    res:=[]int{}  // 右视结果
    if root==nil{
        return res
    }

    currentLevel:=[]*TreeNode{root}
    
    // 处理当前层
    for len(currentLevel)>0{
        // 当前层的值
        levalVal:=[]int{}
        // 下一层节点
        nextLevelNode:=[]*TreeNode{}

        for _,node:=range currentLevel{
            levalVal=append(levalVal,node.Val)
            
            if node.Left!=nil{
                nextLevelNode=append(nextLevelNode,node.Left)
            }
            if node.Right!=nil{
                nextLevelNode=append(nextLevelNode,node.Right)
            }
        }
        // 将最右边的节点值加入到结果集
        res=append(res,levalVal[len(levalVal)-1])

        // 处理下一层
        currentLevel=nextLevelNode
    }
    return res
}

// 2023.01.03