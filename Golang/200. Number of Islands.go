package leetcode

// 岛屿数量
// https://leetcode-cn.com/problems/number-of-islands/

// 问题本质：二维网格grid就是图的邻接矩阵表示，统计无向图中的连通分量的数量
// 图的抽象
//   节点：每个 '1' 格子就是一个节点
//   边：相邻（上下左右）的 '1' 格子之间存在边
// 连通分量：相互连接的 '1' 组成一个连通分量（岛屿）

// 递归DFS
// 思路：扫描整个二维网格。如果一个位置为 1，则以其为起始节点开始进行深度优先搜索。从一个 1 开始，不断向上下左右延伸，直到碰到 0 或边界。
// 在深度优先搜索的过程中，每个搜索到的 1 都会被重新标记为 0。最终岛屿的数量就是进行深度优先搜索的次数。
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// 遍历图：从坐标(r, c)开始，找到它所在的整个岛屿（连通分量）
	var dfs func(int, int)
	dfs = func(r, c int) {
		// 判断坐标(r,c)是否在网格中
		if !(0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])) || grid[r][c] != '1' {
			return
		}

		// 为了避免重复计数和重复遍历，将这个岛屿的所有陆地标记为已访问（避免了额外的visited数组）
		grid[r][c] = '2'

		// 递归调用，遍历上下左右
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	islandNum := 0
	// 遍历所有节点，寻找连通分量
	for i := 0; i < len(grid); i++ { // 行
		for j := 0; j < len(grid[0]); j++ { // 列
			// 找到了一个未被发现的新岛屿的起点，以其为起始节点开始进行深度优先搜索
			if grid[i][j] == '1' {
				dfs(i, j)
				islandNum++ // 岛屿数量加一
			}
		}
	}
	return islandNum
}

// 队列BFS
// 思路：扫描数组，如果一个节点包含1，则以其为根节点启动广度优先搜索。将其放入队列中，并将值设为0以标记为访问过该节点。迭代地搜索队列中的每一个节点，直到队列为空
func numIslandsBFS(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	nr, nc := len(grid), len(grid[0])
	islandNum := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == 1 {
				islandNum++
				grid[r][c] = 2
				queue := [][]int{{r, c}}
				for len(queue) > 0 {
					rc := queue[0]
					queue = queue[1:]
					row, column := rc[0], rc[1]
					if row-1 >= 0 && grid[row-1][column] == 1 {
						queue = append(queue, []int{row - 1, column})
						grid[row-1][column] = 0
					}
					if row+1 < nr && grid[row+1][column] == 1 {
						queue = append(queue, []int{row + 1, column})
						grid[row+1][column] = 0
					}
					if column-1 >= 0 && grid[row][column-1] == 1 {
						queue = append(queue, []int{row, column - 1})
						grid[row][column-1] = 0
					}
					if column+1 < nc && grid[row][column+1] == 1 {
						queue = append(queue, []int{row, column + 1})
						grid[row][column+1] = 0
					}
				}
			}
		}
	}
	return islandNum
}

// UnionFind
