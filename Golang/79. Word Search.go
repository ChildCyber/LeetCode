package leetcode

// 单词搜索
// https://leetcode-cn.com/problems/word-search/

// 回溯
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// 遍历整个网格，寻找起点
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果找到起点，并且从该点开始能找到完整单词，返回true
			if board[i][j] == word[0] && backtrack79(board, word, visited, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func backtrack79(board [][]byte, word string, visited [][]bool, i, j, index int) bool {
	// 如果已经匹配完所有字符，返回true
	if index == len(word) {
		return true
	}

	// 检查边界条件、是否已访问、是否匹配当前字符
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		visited[i][j] || board[i][j] != word[index] {
		return false
	}

	// 标记当前位置为已访问
	visited[i][j] = true

	// 尝试四个方向：上、右、下、左
	if backtrack79(board, word, visited, i-1, j, index+1) || // 上
		backtrack79(board, word, visited, i, j+1, index+1) || // 右
		backtrack79(board, word, visited, i+1, j, index+1) || // 下
		backtrack79(board, word, visited, i, j-1, index+1) { // 左
		return true
	}

	// 回溯：撤销选择
	visited[i][j] = false

	return false
}

// 回溯-闭包
func exist1(board [][]byte, word string) bool {
	rows := len(board)
	if rows == 0 {
		return false
	}
	cols := len(board[0])

	// 防止同一个格子在一条路径中被重复访问
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	// r：当前访问的行号（row）
	// c：当前访问的列号（column）
	// index：当前要匹配 word 的第几个字符
	// 从 board[r][c] 出发，能不能从 word[index:] 继续匹配下去
	var dfs func(r, c, index int) bool
	dfs = func(r, c, index int) bool {
		// 匹配完所有字符，返回 true
		if index == len(word) {
			return true
		}
		// 越界或字符不匹配或已访问过，剪枝
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}
		if board[r][c] != word[index] || visited[r][c] {
			return false
		}

		// 标记当前格子
		visited[r][c] = true
		// 往四个方向探索
		found := dfs(r+1, c, index+1) ||
			dfs(r-1, c, index+1) ||
			dfs(r, c+1, index+1) ||
			dfs(r, c-1, index+1)
		// 回溯：取消标记
		visited[r][c] = false

		return found
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
