package leetcode

// 矩阵置零
// https://leetcode.cn/problems/set-matrix-zeroes/
// 如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0

// 创建副本进行标记
// 空间复杂度：O(mn)
func setZeroesCopy(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return
	}

	// 创建副本矩阵
	copyMatrix := make([][]int, m)
	for i := range copyMatrix {
		copyMatrix[i] = make([]int, n)
		// 复制原始数据
		for j := 0; j < n; j++ {
			copyMatrix[i][j] = matrix[i][j]
		}
	}

	// 遍历副本矩阵，在原始矩阵中设置零
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if copyMatrix[i][j] == 0 {
				// 将原始矩阵的第 i 行设为 0
				for k := 0; k < n; k++ {
					matrix[i][k] = 0
				}
				// 将原始矩阵的第 j 列设为 0
				for l := 0; l < m; l++ {
					matrix[l][j] = 0
				}
			}
		}
	}
}

// 使用标记数组
// 空间复杂度：O(m+n)
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	// 使用两个数组记录需要置零的行和列
	rowZero := make([]bool, m) // 标记哪些行需要置零
	colZero := make([]bool, n) // 标记哪些列需要置零

	// 第一次遍历：标记需要置零的行和列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowZero[i] = true // i行需要都标为0
				colZero[j] = true // j列需要都标为0
			}
		}
	}
	// 第二次遍历：根据标记置零
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果当前元素所在的行或列被标记需要置零，则设置该元素为0
			if rowZero[i] || colZero[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// 使用两个标记变量：用matrix第一行和第一列记录该行该列是否有0
// 空间复杂度：O(1)
func setZeroes1(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])

	// 步骤1：检查第一行和第一列是否有零
	// 记录第一行和第一列是否原本包含0
	rowZero, colZero := false, false
	// 检查第一行是否有零
	for i := 0; i < cols; i++ {
		if matrix[0][i] == 0 {
			rowZero = true
			break
		}
	}
	// 检查第一列是否有零
	for j := 0; j < rows; j++ {
		if matrix[j][0] == 0 {
			colZero = true
			break
		}
	}

	// 步骤2：使用第一行和第一列作为标记空间
	// 遍历除第一行和第一列外的所有元素
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				// 使用第一行和第一列记录需要置零的行和列
				matrix[i][0] = 0 // 标记第i行需要置零
				matrix[0][j] = 0 // 标记第j列需要置零
			}
		}
	}

	// 步骤3：根据标记置零（处理除第一行和第一列外的部分）
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			// 如果当前元素所在的行或列被标记为需要置零，则置零
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 步骤4：处理第一行和第一列
	// 如果第一行需要置零，则置零第一行
	if rowZero {
		for i := 0; i < cols; i++ {
			matrix[0][i] = 0
		}
	}
	// 如果第一列需要置零，则置零第一列
	if colZero {
		for j := 0; j < rows; j++ {
			matrix[j][0] = 0
		}
	}
}
