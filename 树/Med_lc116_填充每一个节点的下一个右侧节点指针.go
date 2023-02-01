package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// 方法一：按层序遍历树解法
func connect1(root *Node) *Node {
	if root==nil{
        return nil
    }

    currentLevelNode:=[]*Node{root}
    for len(currentLevelNode)!=0{
        nextLevelNode:=[]*Node{}
        for i,node :=range currentLevelNode{
            if i+1<len(currentLevelNode){
                node.Next=currentLevelNode[i+1]
            }

            if node.Left!=nil{
                nextLevelNode=append(nextLevelNode,node.Left)
            }
            if node.Right!=nil{
                nextLevelNode=append(nextLevelNode,node.Right)
            }
        }

        currentLevelNode=nextLevelNode
    }

    return root
}

// 方法二：
func connect(root *Node) *Node {
    if root==nil{
        return nil
    }

    for leftMost:=root;leftMost.Left!=nil;leftMost=leftMost.Left{
        for node:=leftMost;node!=nil;node=node.Next{
            node.Left.Next=node.Right

            if node.Next!=nil{
                node.Right.Next=node.Next.Left
            }
        }
    }
    return root
}