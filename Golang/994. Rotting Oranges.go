package leetcode

// 腐烂的橘子
// https://leetcode.cn/problems/rotting-oranges/

// 问题本质：多源最短路径问题
// 图的抽象：
//   顶点：网格中的每一个单元格（0, 1, 或 2）
//   边：相邻的单元格之间有边（上下左右）
// 问题建模：
//   多源：所有的初始腐烂橘子都是感染的起点，它们会同时向四周传播腐烂
//   最短路径：从任意一个腐烂橘子到最远新鲜橘子的最小距离（也就是最小分钟数）

// BFS
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	queue := [][2]int{} // 存储腐烂橘子的位置
	freshCount := 0     // 新鲜橘子计数

	// 初始化：统计新鲜橘子，将腐烂橘子加入队列
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				freshCount++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	// 如果没有新鲜橘子，直接返回0
	if freshCount == 0 {
		return 0
	}

	// 方向数组：上下左右
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0

	// BFS开始
	for len(queue) > 0 {
		levelSize := len(queue)
		infected := false // 标记本轮是否感染了新橘子

		// 处理当前分钟的所有腐烂橘子
		for i := 0; i < levelSize; i++ {
			// 出队，取出一个腐烂橘子
			current := queue[0]
			queue = queue[1:]

			// 向四个方向传播
			row, col := current[0], current[1]
			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]

				// 检查新位置是否有效且是新鲜橘子
				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == 1 {
					// 感染这个新鲜橘子
					grid[newRow][newCol] = 2
					freshCount--
					// 新腐烂橘子加入队列
					queue = append(queue, [2]int{newRow, newCol})
					infected = true
				}
			}
		}

		// 如果本轮感染了新橘子，时间增加
		if infected {
			minutes++
		}
	}

	// 检查是否所有新鲜橘子都被感染
	if freshCount == 0 {
		return minutes // 所有橘子都腐烂了
	}
	return -1 // 有橘子无法被感染
}

// DFS
// 模拟每一分钟的感染事件
func orangesRottingDFS(grid [][]int) int {
	// 初始化时间分钟数
	minutes := 0
	// 调用递归函数
	return dfs(grid, minutes)
}

func dfs(grid [][]int, minutes int) int {
	// 检查是否还有新鲜橘子
	if noFreshOranges(grid) {
		return minutes
	}

	// 标记本分钟会感染的新鲜橘子
	changed := false
	// 临时标记：先记录哪些位置会被感染
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				// 感染上下左右的新鲜橘子，临时标记为3
				infectNeighbors(grid, i, j)
			}
		}
	}

	// 正式更新网格：将临时标记3的橘子变为2
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 3 {
				grid[i][j] = 2
				changed = true
			}
		}
	}

	// 如果本分钟没有感染任何新橘子，但还有新鲜橘子，返回-1
	if !changed && hasFreshOranges(grid) {
		return -1
	}

	// 递归调用下一分钟
	return dfs(grid, minutes+1)
}

// 辅助函数：感染相邻新鲜橘子，临时标记为3
func infectNeighbors(grid [][]int, i, j int) {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) && grid[ni][nj] == 1 {
			grid[ni][nj] = 3 // 临时标记，避免同一分钟内重复感染
		}
	}
}

// 辅助函数：检查是否还有新鲜橘子
func noFreshOranges(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return false
			}
		}
	}
	return true
}

// 辅助函数：检查是否有新鲜橘子
func hasFreshOranges(grid [][]int) bool {
	return !noFreshOranges(grid)
}
