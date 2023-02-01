package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 自底向上

// 两个直接相连的房子不能同时选  ====> 选父一定不能选子，不选父可以考虑选子
func rob(root *TreeNode) int {
    if root==nil{    
        return 0
    }else if root.Left==nil&&root.Right==nil{   // 整棵树只有一个节点
        return root.Val
    }

    var ymap=map[*TreeNode]int{}  // 选此节点的 累计至此的最大总金额
    var nmap=map[*TreeNode]int{}  // 不选此节点的 累计至此的最大总金额

    var dfs func(*TreeNode)
    dfs=func(node *TreeNode){
        if node==nil{
            return
        }
        
        dfs(node.Left)
        dfs(node.Right)
        // 选此节点,则不能选左右子节点, 累计至此节点上的最大总金额则为此节点上的金额+左右子节点的累计金额
        ymap[node]=node.Val+nmap[node.Left]+nmap[node.Right]

        // 不选此节点 则考虑要不要选子节点， 则累计至此节点上的最大总金额 来自其左右子节点的累计金额
        nmap[node]=max(ymap[node.Left],nmap[node.Left])+max(ymap[node.Right],nmap[node.Right])
    }
    dfs(root)
    return max(ymap[root],nmap[root])
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}