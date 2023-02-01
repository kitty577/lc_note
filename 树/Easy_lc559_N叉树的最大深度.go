package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// bfs
// func maxDepth(root *Node) int {
//     if root==nil{
//         return 0
//     }

//     currentLevelNode:=[]*Node{root}
//     ans:=0

//     for len(currentLevelNode)!=0{
//         nextLevelNode:=[]*Node{}
//         for _,node:=range currentLevelNode{
//             if len(node.Children)!=0{
//                 nextLevelNode=append(nextLevelNode,node.Children...)
//             }
//         }
//         currentLevelNode=nextLevelNode
//         ans++
//     }

//     return ans
// }

// dfs
func maxDepth(root *Node) int {
    if root==nil{
        return 0
    }
    maxChildDep:=0
    for _,child:=range root.Children{
        maxChildDep=max(maxChildDep,maxDepth(child))
    }

    return maxChildDep+1
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
