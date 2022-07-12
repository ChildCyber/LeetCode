package leetcode

// 被围绕的区域
// https://leetcode.cn/problems/surrounded-regions
// DFS
func solve(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1 { // 边界，行为第一行或最后一行，列为第一列或最后一列
				if board[i][j] == 'O' { // 边界上的 'O'
					dfs130(i, j, board)
				}
			}
		}
	}

	// 遍历矩阵，对于每一个字母：
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '*' { // 如果该字母被标记过，则该字母为没有被字母 X 包围的字母 O，将其还原为字母 O
				board[i][j] = 'O'
			} else if board[i][j] == 'O' { // 如果该字母没有被标记过，则该字母为被字母 X 包围的字母 O，将其修改为字母 X
				board[i][j] = 'X'
			}
		}
	}
}

func dfs130(i, j int, board [][]byte) {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 { // 判断是否越界
		return
	}

	// 将边缘上的 'O' 标记成 '*'
	if board[i][j] == 'O' {
		board[i][j] = '*'
		for k := 0; k < 4; k++ { // 对于每一个边界上的 O，以它为起点，标记所有与它直接或间接相连的字母 O
			dfs130(i+dir[k][0], j+dir[k][1], board)
		}
	}
}
