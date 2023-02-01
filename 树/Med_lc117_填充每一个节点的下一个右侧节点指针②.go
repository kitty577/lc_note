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

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	currentLevelNode := []*Node{root}
	for len(currentLevelNode) != 0 {
		nextLevelNode := []*Node{}

		for i, node := range currentLevelNode {
			if i+1 < len(currentLevelNode) {
				node.Next = currentLevelNode[i+1]
			}
			if node.Left != nil {
				nextLevelNode = append(nextLevelNode, node.Left)
			}
			if node.Right != nil {
				nextLevelNode = append(nextLevelNode, node.Right)
			}
		}
		currentLevelNode = nextLevelNode
	}
	return root
}
