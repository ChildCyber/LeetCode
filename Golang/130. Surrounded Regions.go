package leetcode

// 被围绕的区域
// https://leetcode.cn/problems/surrounded-regions/

// 问题本质：图连通性问题，考察的是在边界条件下的连通分量识别
// board可以建模为一个图：
//   节点：每个格子就是一个节点
//   边：相邻（上下左右）的格子之间存在边
//   特殊节点：边界上的 'O' 节点
//   连通性：与边界相连的 'O' 形成一个连通分量

// 递归DFS
func solve(board [][]byte) {
	// 逆向思维：与其去找“应该被翻转的”，不如先标记“安全的”
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	// 从边界开始DFS遍历，标记所有与边界相连的'O'
	var dfs func(int, int)
	dfs = func(r, c int) {
		// 检查边界条件和当前字符
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != 'O' {
			return
		}
		// 'O'标记为临时字符'#'，表示已访问且是安全的
		board[r][c] = '#'
		// 向四个方向递归搜索，遍历上下左右
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	// 1. 标记所有边界上的'O'及其相连区域
	// 遍历第一列和最后一列
	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][cols-1] == 'O' {
			dfs(i, cols-1)
		}
	}
	// 遍历第一行和最后一行
	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[rows-1][j] == 'O' {
			dfs(rows-1, j)
		}
	}

	// 2. 第二次遍历：处理标记
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				// 被包围的'O'，翻转为'X'
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				// 安全的'O'，恢复为'O'
				board[i][j] = 'O'
			}
			// 'X' 保持不变
		}
	}
}
