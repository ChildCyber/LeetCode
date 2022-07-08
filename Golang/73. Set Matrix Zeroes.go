package leetcode

// 矩阵置零
// https://leetcode.cn/problems/set-matrix-zeroes/
// 如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0

// 使用标记数组
func setZeroes(matrix [][]int) {
	// 记录每一行是否有零出现
	row := make([]bool, len(matrix))
	// 记录每一列是否有零出现
	col := make([]bool, len(matrix[0]))

	// 第一遍记录哪些行、列有0
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 { // matrix[i][j]==0
				row[i] = true // i行需要都标为0
				col[j] = true // j列需要都标为0
			}
		}
	}

	// 第二遍置0
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] { // i行需要标为0或j列需要标为0
				r[j] = 0 // 第i行的每个元素都设为0
			}
		}
	}
}

// 使用两个标记变量
func setZeroes1(matrix [][]int) {
	// 用matrix第一行和第一列记录该行该列是否有0
	n, m := len(matrix), len(matrix[0])
	// 记录第一行和第一列是否原本包含0
	row0, col0 := false, false

	// 预处理出两个标记变量
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}

	// 使用其他行与列去处理第一行与第一列
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 使用第一行与第一列去更新其他行与列
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 使用两个标记变量更新第一行与第一列
	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}
