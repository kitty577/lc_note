package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归版本
func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1==nil{
        return root2
    }

    if root2==nil{
        return root1
    }

    root1.Val+=root2.Val

    root1.Left=mergeTrees1(root1.Left,root2.Left)
    root1.Right=mergeTrees1(root1.Right,root2.Right)   

    return root1
}


// 迭代版本

// 用队列记录树节点
func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1==nil{
        return root2
    }

    if root2==nil{
        return root1
    }

    queue:=[]*TreeNode{root1,root2}

    for len(queue)>0{
        node1:=queue[0]
        queue=queue[1:]
        node2:=queue[0]
        queue=queue[1:]

        node1.Val+=node2.Val

        // 两棵树都有左子树
        if node1.Left!=nil&&node2.Left!=nil{
            queue=append(queue,node1.Left)
            queue=append(queue,node2.Left)
        }

        // 两棵树都有右子树
        if node1.Right!=nil&&node2.Right!=nil{
            queue=append(queue,node1.Right)
            queue=append(queue,node2.Right)
        }

        if node1.Left==nil{
            node1.Left=node2.Left
        }

        if node1.Right==nil{
            node1.Right=node2.Right
        }
    }

    return root1
}

// 2023.02.01