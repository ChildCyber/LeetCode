package leetcode

// 迷宫中离入口最近的出口
// https://leetcode.cn/problems/nearest-exit-from-entrance-in-maze/

// BFS
func nearestExit(maze [][]byte, entrance []int) int {
	rows := len(maze)
	if rows == 0 {
		return -1
	}
	cols := len(maze[0])

	// 方向数组：上、右、下、左
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// 初始化队列和访问标记
	queue := [][]int{entrance} // 将入口位置加入队列
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	visited[entrance[0]][entrance[1]] = true // 标记入口为已访问

	steps := 0

	for len(queue) > 0 {
		// 当前层的节点数
		levelSize := len(queue)

		// 遍历当前层的所有节点
		for i := 0; i < levelSize; i++ {
			// 出队
			current := queue[0]
			queue = queue[1:]
			r, c := current[0], current[1]

			// 检查是否是出口（边界且不是入口）
			if (r == 0 || r == rows-1 || c == 0 || c == cols-1) &&
				!(r == entrance[0] && c == entrance[1]) {
				return steps
			}

			// 向四个方向探索
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]

				// 检查新位置是否合法
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
					maze[nr][nc] == '.' && !visited[nr][nc] {

					// 标记为已访问并加入队列
					visited[nr][nc] = true
					queue = append(queue, []int{nr, nc})
				}
			}
		}

		// 当前层处理完毕，步数加1
		steps++
	}

	// 没有找到出口
	return -1
}
