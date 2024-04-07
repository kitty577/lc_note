package main

// 根据题意 二分图中边的两个端点分别来自两个集合
// 遍历结点，对于一个结点from 那么它指向的To结点，与from节点是不同集合的。同理，To指向的Three结点， 与from节点必须是同一个集合的。
// 实现上，可以将结点标志，Red Green
const (
	UnColored = 0
	Red       = 1
	Green     = 2
)

func isBipartite(graph [][]int) bool {
	var fit = true
	length := len(graph)
	colorFlag := make([]int, length)

	var dfs func(node, flag int)
	dfs = func(node, flag int) {
		colorFlag[node] = flag
		toFlag := Red
		if flag == Red {
			toFlag = Green
		} // 目标节点和起始节点不同标记，标明来自两个不同的集合

		for _, neighbor := range graph[node] {
			if colorFlag[neighbor] == UnColored {
				dfs(neighbor, toFlag)
				if !fit {
					return
				}
			} else if colorFlag[neighbor] != toFlag {
				fit = false
				return
			}
		}
	}

	for i := 0; i < length && fit; i++ {
		if colorFlag[i] == UnColored {
			dfs(i, Red)
		}
	}

	return fit
}
