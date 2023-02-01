package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 迭代： 每次记录最左边的节点的值 随着循环 更新值
func findBottomLeftValue1(root *TreeNode) int {
    var gradation int
    queue := list.New()
    
    queue.PushBack(root)
    for queue.Len() > 0 {
        length := queue.Len()
        for i := 0; i < length; i++ {
            node := queue.Remove(queue.Front()).(*TreeNode)
            if i == 0 {
                gradation = node.Val
            }
            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
        }
    }
    return gradation
}

// 递归
// 由于要求的是 最底层、最左边 因此需要记录树的高度 
func findBottomLeftValue(root *TreeNode) int {
    depth,res:=0,0
    var dfs func(*TreeNode,int)
    dfs=func(node *TreeNode,dep int){
        if node==nil{
            return 
        }
        if node.Left==nil&&node.Right==nil&&depth<dep{
            depth=dep
            res=node.Val
        }

        dfs(node.Left,dep+1)
        dfs(node.Right,dep+1)
    }
    
    dfs(root,1)
    return res
}

