package leetcode

// 螺旋矩阵
// https://leetcode-cn.com/problems/spiral-matrix/

// 模拟
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total       = rows * columns
		order       = make([]int, total)
		row, column = 0, 0
		// 四个方向：右下左上
		directions     = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}

// 按层模拟
// 将矩阵看成若干层，首先输出最外层的元素，其次输出次外层的元素，直到输出最内层的元素
// 时间复杂度：O(m×n)
// 空间复杂度：O(1)
func spiralOrderLayer(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 { // 如果 matrix 为空，直接访问 matrix[0] 会 panic
		return []int{}
	}

	rows, columns := len(matrix), len(matrix[0])
	order := make([]int, 0, rows*columns)
	// 定义四个边界
	left, right, top, bottom := 0, columns-1, 0, rows-1

	for left <= right && top <= bottom { // 只要有一个边界条件不满足，就意味着整个矩阵已经遍历完毕
		// 从左到右遍历上边界
		for col := left; col <= right; col++ {
			order = append(order, matrix[top][col])
		}
		top++ // 上边界下移

		// 从上到下遍历右边界
		for row := top; row <= bottom; row++ {
			order = append(order, matrix[row][right])
		}
		right-- // 右边界左移

		// 检查是否还有行需要处理
		if top <= bottom {
			// 从右到左遍历下边界
			for col := right; col >= left; col-- {
				order = append(order, matrix[bottom][col])
			}
			bottom-- // 下边界上移
		}

		// 检查是否还有列需要处理
		if left <= right {
			// 从下到上遍历左边界
			for row := bottom; row >= top; row-- {
				order = append(order, matrix[row][left])
			}
			left++ // 左边界右移
		}
	}
	return order
}
