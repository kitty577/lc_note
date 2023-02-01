package main

// 注意 这题与岛屿数量不同 岛屿数量遍历的是坐标位置 对坐标位置进行dfs
// 这题是遍历城市，对城市进行dfs

///===> 原因：间接相连这个条件

func findCircleNum(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !vis[to] {
				dfs(to)
			}
		}
	}
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return
}
