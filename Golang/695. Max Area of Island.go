package leetcode

// 岛屿的最大面积
// https://leetcode-cn.com/problems/max-area-of-island/
// DFS
func maxAreaOfIsland695(grid [][]int) int {
	ans := 0
	// 遍历所有元素
	for r := 0; r < len(grid); r++ { // 行
		for c := 0; c < len(grid[0]); c++ { // 列
			if grid[r][c] == 1 { // 当前元素为岛屿
				a := area(grid, r, c)
				ans = max(ans, a)
			}
		}
	}
	return ans
}

func isInGrid(grid [][]int, r, c int) bool { // 判断是否越界
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func area(grid [][]int, r, c int) int { // 返回以gird[r][c]为起点的岛屿的面积
	if !isInGrid(grid, r, c) {
		return 0
	}

	if grid[r][c] != 1 {
		return 0
	}

	grid[r][c] = 2 // 当前岛屿标记为已访问

	// 岛屿自身面积加上 上下左右 岛屿的面积
	return 1 +
		area(grid, r-1, c) +
		area(grid, r+1, c) +
		area(grid, r, c-1) +
		area(grid, r, c+1)
}
