package leetcode

// N皇后
// https://leetcode-cn.com/problems/n-queens/
// 回溯
var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{} // 返回结果
	queens := make([]int, n) // 路径：存储皇后，小于row的那些行都已经成功放置了皇后
	for i := 0; i < n; i++ { // 填充数组queens[]中的每个元素都是-1
		queens[i] = -1
	}
	columns := map[int]bool{} // 记录当前元素所在的列，并将其所在的列标记为不可放元素
	// diagonals1：记录当前元素所在的左对角线，并将他所在的左对角线标记为不可放元素
	// diagonals2：记录当前元素所在的右对角线，并将他所在的右对角线标记为不可放元素
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, diagonals1, diagonals2)
	return solutions
}

func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	// 遍历到最后一行，就说明存在一个解法，将皇后的位置存放入结果中
	if row == n {
		board := generateBoard(queens, n)    // 将当前的N行N列中的元素所在的位置结果，转换格式
		solutions = append(solutions, board) // 将符合条件的结果添加进返回结果集中
		return
	}
	// 遍历所有行
	for i := 0; i < n; i++ {
		if columns[i] { // 当前行元素所在的列，都不可放元素
			continue
		}
		// 去除左对角线上的所有元素
		// row 表示行，i表示列
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		// 去除右对角线上的元素
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		// 经过上面的三次排除，就可以找到元素在当前行的哪一列的位置
		// 选第一行的第几列，也可以叫单元格所在的位置
		queens[row] = i
		// 把选中的单元格加入到，去除列的集合中
		// 用来给下一行的元素所在的列作为排除条件判断
		columns[i] = true
		// 把选中的单元格加入到，去除左对角线的集合中
		// 把选中的单元格加入到，去除右对角线的集合中
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		// 递归遍历下一行
		backtrack(queens, n, row+1, columns, diagonals1, diagonals2)
		// 回溯操作
		// 剪枝操作
		queens[row] = -1
		// 将当前列和左对角线和右对角线的元素都删除，避免重复遍历
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

// 转换格式
func generateBoard(queens []int, n int) []string {
	board := []string{}
	// 遍历所有行
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		// 将当前行所在的列的，位置置为Q
		row[queens[i]] = 'Q'
		// 将当前结果添加进结果集中
		board = append(board, string(row))
	}
	return board
}
