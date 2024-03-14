package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 第一种 中序遍历，找到节点的下一个即可。 由于只是需要一个，并不需要全部的中序遍历结果，设法只记录一个preNode
/*
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    listN:=[]*TreeNode{}
    // 将二叉搜索树的中序遍历结果存储在数组中
    var inorder func(node *TreeNode)
    inorder=func(node *TreeNode){
        if node==nil{
            return
        }
        // 左
        inorder(node.Left)

        // 中 处理节点
        listN=append(listN,node)

        // 右
        inorder(node.Right)
    }

    // 二分查找 找到p节点在数组中的下标
    var search func(listN []*TreeNode)int
    search=func(listN []*TreeNode) int{
        left,right:=0,len(listN)-1
        for left<=right{
            mid:=left+(right-left)/2
            if listN[mid].Val==p.Val{
                return mid
            }else if listN[mid].Val>p.Val{
                right=mid-1
            }else{
                left=mid+1
            }
        }
        return left
    }

    inorder(root)
    index:=search(listN)
    if index>=len(listN)-1{
        return nil
    }

    return listN[index+1]
}
*/

// 利用二叉搜索树特性。 如果一个节点有右子树，那么后继节点为右子树的最左的节点；
// 如果没有右子树那么只能往上找父节点的兄弟节点
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ans *TreeNode
	if p.Right != nil {
		ans = p.Right
		for ans.Left != nil {
			ans = ans.Left
		}
		return ans
	}

	node := root
	for node != nil {
		if node.Val > p.Val {
			ans = node
			node = node.Left
		} else if node.Val < p.Val {
			node = node.Right
		} else if node.Val == p.Val {
			return ans
		}
	}
	return nil
}
