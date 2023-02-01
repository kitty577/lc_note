package main

// 对于图中的任意一个节点，它在搜索的过程中有三种状态，即：

// 		「未搜索」：我们还没有搜索到这个节点；

// 		「搜索中」：我们搜索过这个节点，但还没有回溯到该节点，即该节点还没有入栈，还有相邻的节点没有搜索完成）；

// 		「已完成」：我们搜索过并且回溯过这个节点，即该节点已经入栈，并且所有该节点的相邻节点都出现在栈的更底部的位置，满足拓扑排序的要求。

// 通过上述的三种状态，我们就可以给出使用深度优先搜索得到拓扑排序的算法流程，在每一轮的搜索搜索开始时，我们任取一个「未搜索」的节点开始进行深度优先搜索。

// 我们将当前搜索的节点 u 标记为「搜索中」，遍历该节点的每一个相邻节点 v：

// 		如果 v 为「未搜索」，那么我们开始搜索 v，待搜索完成回溯到 u；

// 		如果 v 为「搜索中」，那么我们就找到了图中的一个环，因此是不存在拓扑排序的；

// 		如果 v 为「已完成」，那么说明 v 已经在栈中了，而 u 还不在栈中，因此 u 无论何时入栈都不会影响到 (u, v) 之前的拓扑关系，以及不用进行任何操作。

// 当 u 的所有相邻节点都为「已完成」时，我们将 u 放入栈中，并将其标记为「已完成」。

// 在整个深度优先搜索的过程结束后，如果我们没有找到图中的环，那么栈中存储这所有的 n 个节点，从栈顶到栈底的顺序即为一种拓扑排序。

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges  = make([][]int, numCourses) // 存储有向图
		visitd = make([]int, numCourses)   // 存储每个节点的状态 0-未搜索，1-搜索中，2-搜索完成
		result = make([]int, 0)            // 数组模拟栈，栈顶的是最基础的结点
		valid  = true
		dfs    func(int)
	)

	// 统计边
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	dfs = func(u int) {
		visitd[u] = 1
		for _, u := range edges[u] {
			if visitd[u] == 0 {
				dfs(u)
				if !valid {
					return
				}
			} else if visitd[u] == 1 { // 说明别的结点也在搜索这个点==》形成了环
				valid = false
				return
			}
		}
		// 搜索完了这个节点
		visitd[u] = 2
		result = append(result, u)
	}

	for i := 0; i < numCourses; i++ {
		if visitd[i] == 0 {
			dfs(i)
		}

	}

	if !valid {
		return []int{}
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}

// 2022.12.26
