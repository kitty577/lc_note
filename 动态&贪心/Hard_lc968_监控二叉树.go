package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 把摄像头放在叶子节点的父节点上 这样监控的覆盖节点多
// 那么就要从叶子节点往上走 找父节点 ===》后序遍历树 左右中
// 每个摄像头可以覆盖本身、上一层、下一层
// 那么可以从低向上 叶子节点的父节点放摄像头、然后隔两个节点（两层）再放一个摄像头，直至遍历到根节点

// 那怎么做到隔两层呢？ dp记录节点的覆盖状况  0：该节点未安装摄像头且未被覆盖  1：该节点本身安装摄像头 2：该节点未安装摄像头但被覆盖

// 递归 遍历到空节点 空节点的状态应该设为2
func minCameraCover(root *TreeNode) int {
	ans := 0

	var travel func(cur *TreeNode) int
	travel = func(cur *TreeNode) int {
		if cur == nil {
			return 2
		}
		// 后序遍历
		left := travel(cur.Left)   // 左
		right := travel(cur.Right) // 右
		// 中 处理

		//  左右子节点都被覆盖
		if left == 2 && right == 2 {
			return 0
		}

		// 左右子节点有未被覆盖的
		if left == 0 || right == 0 {
			ans++
			return 1
		}

		// 左右节点 其中有装摄像头的
		if left == 1 || right == 1 {
			return 2
		}
		return 0
	}

	r := travel(root)
	// 如果遍历完所有节点 而头结点没有被覆盖到 那么在头结点上装一个
	if r == 0 {
		ans++
	}

	return ans

}
