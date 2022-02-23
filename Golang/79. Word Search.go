package leetcode

// 单词搜索
// https://leetcode-cn.com/problems/word-search/
// 回溯
// 在地图上的任意一个起点开始，向 4 个方向分别 DFS 搜索，直到所有的单词字母都找到了就输出 true，否则输出 false
var dir = [][]int{ // 上下左右
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board)) // 标记是否访问过
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board { // 行
		for j := range v { // 列
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool { // 判断是否越界
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}

	if board[x][y] == word[index] { // 剪枝：当前字符匹配
		visited[x][y] = true     // 标记为已访问
		for i := 0; i < 4; i++ { // 朝四个方向搜索
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false // 回溯，标记为未访问
	}
	return false
}
