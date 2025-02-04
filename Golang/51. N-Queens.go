package leetcode

// N皇后
// https://leetcode-cn.com/problems/n-queens/

// 回溯
func solveNQueens1(n int) [][]string {
	// 搜索策略：
	// 1. 按行放置皇后（一行一行来）
	// 2. 在每一行中，尝试所有可能的列位置
	// 3. 如果当前位置安全，就放置皇后并继续下一行
	// 4. 如果遇到冲突，就回退到上一步，尝试其他位置
	var res [][]string
	board := make([][]byte, n) // 记录当前棋盘状态
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	cols := make(map[int]bool)  // 记录已占用的列，每行皇后所在的列位置
	diag1 := make(map[int]bool) // 主对角线 row-col
	diag2 := make(map[int]bool) // 副对角线 row+col

	var backtrack func(row int)
	backtrack = func(row int) {
		// 结束条件：row == n，找到一个解
		if row == n {
			solution := make([]string, n)
			for i := 0; i < n; i++ {
				solution[i] = string(board[i])
			}
			res = append(res, solution)
			return
		}

		// 遍历每一列，尝试放皇后
		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue // 冲突，不合法
			}

			// 做选择
			board[row][col] = 'Q'
			cols[col] = true
			diag1[row-col] = true
			diag2[row+col] = true

			// 进入下一行
			backtrack(row + 1)

			// 撤销选择
			board[row][col] = '.'
			cols[col] = false
			diag1[row-col] = false
			diag2[row+col] = false
		}
	}

	backtrack(0)
	return res
}
