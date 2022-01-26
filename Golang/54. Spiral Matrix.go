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
		// 朝右，朝下，朝左，朝上
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
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
func spiralOrder1(matrix [][]int) []int {
	// 提前算出一共多少个元素,一圈一圈地遍历矩阵,停止条件就是遍历了所有元素(count == sum)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns = len(matrix), len(matrix[0])
		order         = make([]int, rows*columns)
		index         = 0
		// top、left、right、bottom 分别是剩余区域的上、左、右、下的下标
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom { // 可以相遇
		for column := left; column <= right; column++ { // 从左到右遍历上侧元素，依次为 (top,left) 到 (top,right)，right作为边界，以列作为索引
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ { // 从上到下遍历右侧元素，依次为 (top+1,right) 到 (bottom,right)，bottom作为边界，以行作为索引
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom { // 相遇
			for column := right - 1; column > left; column-- { // 从右到左，left作为边界（不可以等于），以列作为索引
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- { // 从下到上，top作为边界（不可以等于），以行作为索引
				order[index] = matrix[row][left]
				index++
			}
		}
		// 进入到下一层
		left++
		right--
		top++
		bottom--
	}
	return order
}
