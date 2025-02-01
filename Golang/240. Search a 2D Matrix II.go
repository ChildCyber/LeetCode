package leetcode

// 搜索二维矩阵 II
// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/

// 直接查找
func searchMatrix(matrix [][]int, target int) bool {
	// 直接遍历整个矩阵matrix，判断target是否出现
	for _, row := range matrix {
		for _, v := range row {
			if v == target {
				return true
			}
		}
	}
	return false
}

// 对每行进行二分查找
func searchMatrixBS(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for _, row := range matrix {
		low, high := 0, len(matrix[0])-1
		for low <= high {
			mid := low + (high-low)/2
			if row[mid] == target {
				return true
			} else if row[mid] > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return false
}

// Z 字形查找（二分变种）
func searchMatrixZ(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// 从矩阵 matrix 的右上角 (0,n−1) 进行搜索
	row, col := 0, len(matrix[0])-1
	for col >= 0 && row <= len(matrix)-1 {
		if target == matrix[row][col] {
			return true
		}

		// 在每一步的搜索过程中，如果位于位置 (x,y)，希望在以 matrix 的左下角为左下角、以 (x,y) 为右上角的矩阵中进行搜索
		if target > matrix[row][col] { // 所有位于该列的元素都是严格大于target的
			row++
		} else { // 所有位于该行的元素都是严格小于target的
			col--
		}
	}

	// 超出矩阵边界，不存在target
	return false
}
