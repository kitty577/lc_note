package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 与105相似，递归构建
// 前序遍历：根--》左--》右
// 中序遍历：左--》根--》右
// 后序遍历：左--》右--》根

// // 根据后序遍历postorder 从后往前 能确定根节点，再通过此根节点在中序遍历中能确定左子树和右子树的区间范围



// func buildTree(inorder []int, postorder []int) *TreeNode {
//     if len(inorder)==0{
//         return nil
//     }
//     lenp:=len(postorder)
//     root:=&TreeNode{postorder[lenp-1],nil,nil} // TODO: 此写法会报错，out of memory!!!!!
//     // 确定左右子树区间
//     i:=0
//     for i,_=range inorder{
//         if root.Val==inorder[i]{
//             break
//         }
//         i++
//     }

//     index:=len(inorder[:i])+1

//     root.Left=buildTree(inorder[:i],postorder[:index])
//     root.Right=buildTree(inorder[i+1:],postorder[index:len(postorder)])
//     return root
// }

// 改用map记录节点位置

var m map[int]int   // 用map保存中序序列的数值对应的位置
func buildTree(inorder []int, postorder []int) *TreeNode {
    m=make(map[int]int)
    for i,v:=range inorder{
        m[v]=i
    }

    root:=rebuild(inorder,postorder,len(postorder)-1,0,len(inorder)-1)
    return root
}

// rootIdx表示根节点在后序数组中的索引，l, r 表示在中序数组中的前后切分点
func rebuild(inorder []int, postorder []int, rootIdx int, l, r int) *TreeNode {
    // 没有元素
    if l>r{
        return nil
    }
    // 只有一个元素
    if l==r{  
        return &TreeNode{inorder[l],nil,nil}
    }

    // 根据后序数组找到根节点的位置
    rootV:=postorder[rootIdx]
    // 根节点在中序数组中的位置
    rootPos:=m[rootV]

    root:=&TreeNode{rootV,nil,nil}
    root.Left=rebuild(inorder,postorder,rootIdx-(r-rootPos)-1,l,rootPos-1)
    root.Right=rebuild(inorder,postorder,rootIdx-1,rootPos+1,r)
    return root
}