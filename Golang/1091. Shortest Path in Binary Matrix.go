package leetcode

// 二进制矩阵中的最短路径
// https://leetcode.cn/problems/shortest-path-in-binary-matrix/

// 问题本质：在无权无向图（每步代价相同）的网格中寻找最短路径的问题
// 图的抽象：
//   顶点：图中的每一个可通行的单元格
//   边：节点之间的可行走移动
// 问题建模：
//   单源：给定起点左上角 (0, 0)
//   最短路径：单个终点右下角 (N-1, N-1)，需要找到步数最少的路径

// BFS
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	// 检查起点或终点是否为障碍物
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	// 8个移动方向
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	queue := [][3]int{{0, 0, 1}} // [row, col, distance]
	grid[0][0] = 1               // 标记为已访问

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		row, col, dist := cur[0], cur[1], cur[2]

		// 到达终点
		if row == n-1 && col == n-1 {
			return dist
		}

		// 向8个方向探索
		for _, dir := range directions {
			nr, nc := row+dir[0], col+dir[1]

			if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] == 0 {
				grid[nr][nc] = 1 // 标记为已访问
				queue = append(queue, [3]int{nr, nc, dist + 1})
			}
		}
	}

	return -1
}
