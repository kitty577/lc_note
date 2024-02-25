package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root      *TreeNode
	candidate []*TreeNode // 存储未满的叶节点
}

func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	candidate := []*TreeNode{}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}

		if node.Left == nil || node.Right == nil {
			candidate = append(candidate, node)
		}
	}
	return CBTInserter{root, candidate}
}

func (this *CBTInserter) Insert(v int) int {
	// 找未满的叶子节点插入
	// 构建新节点
	newNode := &TreeNode{Val: v}
	node := this.candidate[0] // 从未满的叶节点中取一个出来，将新节点挂在上面
	if node.Left == nil {
		// 优先填左
		node.Left = newNode
	} else {
		node.Right = newNode
		// 该结点的左结点不为空，右节点被新节点填充，因此将该结点中未满结点中移除
		this.candidate = this.candidate[1:]
	}

	// 新节点未满
	this.candidate = append(this.candidate, newNode)
	return node.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
