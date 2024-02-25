package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历
const (
	Spilt = ","
	Null  = "null"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	builder := strings.Builder{}
	q := []*TreeNode{root}
	// 层序遍历
	for len(q) > 0 {
		if builder.Len() > 0 {
			builder.WriteString(Spilt)
		}
		node := q[0]
		q = q[1:]
		s := Null
		if node != nil {
			s = strconv.Itoa(node.Val)
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
		builder.WriteString(s)
	}
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == Null {
		return nil
	}
	nodes := strings.Split(data, Spilt)

	root := parseNode(nodes[0])

	q := []*TreeNode{root}
	i := 1
	for i < len(nodes) {
		// 按层序遍历
		n := q[0]
		q = q[1:]
		n.Left = parseNode(nodes[i])
		n.Right = parseNode(nodes[i+1])
		i += 2
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
	return root
}

func parseNode(s string) *TreeNode {
	if s == Null {
		return nil
	}
	i, _ := strconv.Atoi(s)
	return &TreeNode{
		Val: i,
	}
}
