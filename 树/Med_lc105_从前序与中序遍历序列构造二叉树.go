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

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder)==0{
        return nil
    }

    root:=&TreeNode{preorder[0],nil,nil}

    // 找到左子树和右子树区域
    i:=0
    for i,_=range inorder{
        if preorder[0]==inorder[i]{
            break
        }
        i++
    }
    // 经过以上步骤，i为根节点，i左为左子树，i右为右子树
    root.Left=buildTree(preorder[1:len(inorder[:i])+1],inorder[:i])
    root.Right=buildTree(preorder[len(inorder[:i])+1:],inorder[i+1:])
    return root
}