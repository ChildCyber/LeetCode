package leetcode

// 岛屿的最大面积
// https://leetcode-cn.com/problems/max-area-of-island/
// DFS
func maxAreaOfIsland695(grid [][]int) int {
	res := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				a := area(grid, r, c)
				res = max(res, a)
			}
		}
	}
	return res
}

func isInGrid(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func area(grid [][]int, r, c int) int {
	if !isInGrid(grid, r, c) {
		return 0
	}
	if grid[r][c] != 1 {
		return 0
	}
	grid[r][c] = 2

	// 岛屿自身面积加上 上下左右 岛屿的面积
	return 1 +
		area(grid, r-1, c) +
		area(grid, r+1, c) +
		area(grid, r, c-1) +
		area(grid, r, c+1)
}
