package leetcode

// 搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix/

// 暴力
func searchMatrix74Force(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// Z 字形查找
// 时间复杂度：O(m+n)
func searchMatrix74Z(matrix [][]int, target int) bool {
	row, col := 0, len(matrix[0])-1
	for col >= 0 && row <= len(matrix)-1 {
		if target == matrix[row][col] {
			return true
		}

		if target > matrix[row][col] {
			row++
		} else {
			col--
		}
	}

	return false
}

// 二分
// 时间复杂度：O(log mn)
func searchMatrix74(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m, low, high := len(matrix[0]), 0, len(matrix[0])*len(matrix)-1
	for low <= high {
		mid := low + (high-low)>>1
		// 行列坐标转换
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
