package leetcode

// 岛屿数量
// https://leetcode-cn.com/problems/number-of-islands/
// DFS
// 为了求出岛屿的数量，扫描整个二维网格
// 如果一个位置为 1，则以其为起始节点开始进行深度优先搜索。在深度优先搜索的过程中，每个搜索到的 1 都会被重新标记为 0。最终岛屿的数量就是进行深度优先搜索的次数。
func maxAreaOfIsland(gird [][]int) int {
	islandNum := 0
	// 遍历数组中的每个元素
	for i := 0; i < len(gird); i++ { // 行
		for j := 0; j < len(gird[0]); j++ { // 列
			if gird[i][j] == 1 { // 位置为1，以其为起始节点开始进行深度优先搜索
				dfs(gird, i, j)
				islandNum++ // 岛屿数量加一
			}
		}
	}
	return islandNum
}

func dfs(grid [][]int, i, j int) {
	if !inArea(grid, i, j) {
		return
	}
	grid[i][j] = 2 // 感染，标记为已访问
	// 遍历上下左右
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

// 判断坐标(i,j)是否在网格中
func inArea(grid [][]int, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0]) && grid[i][j] == 1 // 判断是否越界，grid[i][j] == 1是重点
}

// BFS
// 思路：扫描数组，如果一个节点包含1，则以其为根节点启动广度优先搜索。将其放入队列中，并将值设为0以标记为访问过该节点。迭代地搜索队列中的每一个节点，直到队列为空
func maxAreaOfIslandBFS(grid [][]int) int {
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

func dfs1(grid [][]int, i, j int) {
	grid[i][j] = 0
	if 0 <= i-1 && i-1 < len(grid) && 0 <= j && j < len(grid[0]) && grid[i-1][j] == 1 {
		dfs1(grid, i-1, j)
	}
	if 0 <= i+1 && i+1 < len(grid) && 0 <= j && j < len(grid[0]) && grid[i+1][j] == 1 {
		dfs1(grid, i+1, j)
	}
	if 0 <= i && i < len(grid) && 0 <= j-1 && j-1 < len(grid[0]) && grid[i][j-1] == 1 {
		dfs1(grid, i, j-1)
	}
	if 0 <= i && i < len(grid) && 0 <= j+1 && j+1 < len(grid[0]) && grid[i][j+1] == 1 {
		dfs1(grid, i, j+1)
	}
}
