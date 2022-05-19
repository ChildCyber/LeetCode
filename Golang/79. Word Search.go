package leetcode

// 单词搜索
// https://leetcode-cn.com/problems/word-search/
// 回溯
// 在地图上的任意一个起点开始，向 4 个方向分别 DFS 搜索，直到所有的单词字母都找到了就输出 true，否则输出 false
var dir = [][]int{ // 上下左右4个方向
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func exist(board [][]byte, word string) bool {
	// 为了防止重复遍历相同的位置，额外维护一个与 board 等大的 visited 数组，用于标识每个位置是否被访问过。每次遍历相邻位置时，需要跳过已经被访问的位置
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	// 遍历board
	for i, v := range board { // 行
		for j := range v { // 列
			// 对board中的每一个位置 (i,j) 都进行检查：只要有一处返回 true，就说明网格中能够找到相应的单词，否则说明不能找到
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool { // 判断是否越界
	return 0 <= x && x < len(board) && 0 <= y && y < len(board[0])
}

// 判断以 (i,j) 位置为起点出发，能否搜索到单词 word[index..]，其中 word[index..] 表示字符串 word 从第 index 个字符开始的后缀子串
func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	// index：表示搜索到字符串的下标
	// x,y：表示board中的坐标
	if index == len(word)-1 { // 已经搜索到字符串的末尾
		return board[x][y] == word[index] // 判断最后一个字符是否相等
	}

	if board[x][y] == word[index] { // 剪枝：当前字符匹配
		visited[x][y] = true // 标记为已访问
		//defer func() { visited[x][y] = false }()
		for i := 0; i < 4; i++ { // 朝四个方向搜索，下一个字符
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			// 未越界，未访问，且能搜索到单词 word[index+1..]，搜索下一个字符
			if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false // 回溯，标记为未访问
	}
	return false
}
