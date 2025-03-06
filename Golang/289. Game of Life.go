package leetcode

// 生命游戏
// https://leetcode.cn/problems/game-of-life/

// 使用额外数组
func gameOfLifeCopy(board [][]int) {
	m, n := len(board), len(board[0])
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// 复制当前状态
	copyBoard := make([][]int, m)
	for i := range copyBoard {
		copyBoard[i] = make([]int, n)
		copy(copyBoard[i], board[i])
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			live := 0
			// 数活邻居
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && copyBoard[nr][nc] == 1 {
					live++
				}
			}

			// 应用规则
			if copyBoard[r][c] == 1 && (live < 2 || live > 3) {
				board[r][c] = 0
			} else if copyBoard[r][c] == 0 && live == 3 {
				board[r][c] = 1
			}
		}
	}
}

// 原地更新
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// 0：死 → 死
	// 1：活 → 活
	// 2：死 → 活
	// 3：活 → 死
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			live := 0
			// 数活邻居
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n &&
					(board[nr][nc] == 1 || board[nr][nc] == 3) {
					live++
				}
			}

			// 应用规则
			if board[r][c] == 1 && (live < 2 || live > 3) {
				board[r][c] = 3 // 活 -> 死
			}
			if board[r][c] == 0 && live == 3 {
				board[r][c] = 2 // 死 -> 活
			}
		}
	}

	// 第二次遍历更新到下一代
	// 2 → 设为 1
	// 3 → 设为 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == 2 {
				board[r][c] = 1
			} else if board[r][c] == 3 {
				board[r][c] = 0
			}
		}
	}
}
