package leetcode

// 岛屿的最大面积
// https://leetcode-cn.com/problems/max-area-of-island/

var dir695 = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func maxAreaOfIsland695(grid [][]int) int {
	res := 0 // 岛屿最大面积
	for i, row := range grid {
		for j, col := range row {
			if col == 0 {
				continue
			}
			area := areaOfIsland(grid, i, j) // 岛屿面积
			if area > res {
				res = area
			}
		}
	}
	return res
}

func isInGrid(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func areaOfIsland(grid [][]int, x, y int) int {
	if !isInGrid(grid, x, y) || grid[x][y] == 0 {
		return 0
	}

	grid[x][y] = 0           // 标为0
	total := 1               // 多出来的变量
	for i := 0; i < 4; i++ { // 四个方向
		nx := x + dir695[i][0]
		ny := y + dir695[i][1]
		total += areaOfIsland(grid, nx, ny) // +=
	}
	return total
}
