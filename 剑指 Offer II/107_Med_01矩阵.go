package main

// 给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离
// bfs + 扩散标记
// 找到所有的0单元格，放入队列。依次取出0 上下左右扩散，更新距离值，并将扩散到达的位置加入队列中，下一轮扩散。
// 如给定的初始mat：
// 1 1 1
// 1 0 1
// 1 1 0
// 对应的初始的res：
// 0 0 0
// 0 0 0
// 0 0 0

// 对0开始扩散
// _ 1 _
// 1 0 1
// _ 1 0
// 再对1开始扩散：
// 2 1 2
// 1 0 1
// 2 1 0  // 全部扩散完毕，得到结果。
func updateMatrix(mat [][]int) [][]int {
	row, column := len(mat), len(mat[0])
	res := make([][]int, row) // 存储距离
	for i := range res {
		res[i] = make([]int, column)
	}

	visit := make([][]bool, row) // 标记某个点是否扩散过
	for i := range visit {
		visit[i] = make([]bool, column)
	}

	nodes := make([]pos, 0)

	// 首先找到所有的mat数组中的所有0，作为第一次扩散的节点
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				nodes = append(nodes, pos{
					i: i,
					j: j,
				})

				visit[i][j] = true
			}
		}
	}

	// 方向数组： 上、下、左、右
	dirs := []pos{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for len(nodes) != 0 {
		cur := nodes[0]
		nodes = nodes[1:]        // pop
		for _, d := range dirs { // 上下左右四个方位扩散
			newI := cur.i + d.i
			newJ := cur.j + d.j

			// 剪枝（越界、已经扩散过）
			if newI < 0 || newI >= row || newJ < 0 || newJ >= column || visit[newI][newJ] {
				continue
			}
			res[newI][newJ] = res[cur.i][cur.j] + 1
			visit[newI][newJ] = true

			// 当前点加入队列，作为下一轮的扩散起始点
			nodes = append(nodes, pos{
				i: newI,
				j: newJ,
			})
		}
	}

	return res
}

type pos struct {
	i, j int
}
