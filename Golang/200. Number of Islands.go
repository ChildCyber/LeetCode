package leetcode

// 岛屿数量
// https://leetcode-cn.com/problems/number-of-islands/
// DFS
// 为了求出岛屿的数量，扫描整个二维网格
// 如果一个位置为 1，则以其为起始节点开始进行深度优先搜索。在深度优先搜索的过程中，每个搜索到的 1 都会被重新标记为 0。最终岛屿的数量就是进行深度优先搜索的次数。
func maxAreaOfIsland(gird [][]int) int {
	islandNum := 0
	for i := 0; i < len(gird); i++ { // 行
		for j := 0; j < len(gird[0]); j++ { // 列
			if gird[i][j] == 1 { // 位置为1，以其为起始节点开始进行深度优先搜索
				dfs(gird, i, j)
				islandNum++
			}
		}
	}
	return islandNum
}

func dfs(grid [][]int, i, j int) {
	if !inArea(grid, i, j) {
		return
	}
	grid[i][j] = 2 // 感染
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
