package leetcode

// 岛屿的周长
// https://leetcode-cn.com/problems/island-perimeter/
// 迭代
func islandPerimeter(grid [][]int) int {
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if i-1 < 0 || grid[i-1][j] == 0 { // 上
					counter++
				}
				if i+1 >= len(grid) || grid[i+1][j] == 0 { // 下
					counter++
				}
				if j-1 < 0 || grid[i][j-1] == 0 { // 左
					counter++
				}
				if j+1 >= len(grid[0]) || grid[i][j+1] == 0 { // 右
					counter++
				}
			}
		}
	}
	return counter
}

// DFS
func islandPerimeterDFS(grid [][]int) int {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				// 题目限制只有一个岛屿，计算一个即可
				return dfs463(grid, r, c)
			}
		}
	}
	return 0
}

func dfs463(grid [][]int, r, c int) int {
	// 函数因为「坐标 (r, c) 超出网格范围」返回，对应与网格边界相邻的边
	if !inArea(grid, r, c) {
		return 1
	}
	// 函数因为「当前格子是海洋格子」返回，对应海洋格子相邻的边
	if grid[r][c] == 0 {
		return 1
	}
	// 函数因为「当前格子是已遍历的陆地格子」返回，和周长没关系
	if grid[r][c] != 1 {
		return 0
	}
	grid[r][c] = 2
	return dfs463(grid, r-1, c) +
		dfs463(grid, r+1, c) +
		dfs463(grid, r, c-1) +
		dfs463(grid, r, c+1)
}

func isInGrid463(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
