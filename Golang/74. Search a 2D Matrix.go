package leetcode

// 搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix/

// 暴力
func searchMatrix74Brute(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// Z字形查找 同240题
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

// 二分查找
// 时间复杂度：O(log mn)
// 该矩阵可以看作一个完整的有序一维数组，可以直接使用二分查找。关键是如何在二维坐标和一维索引之间进行转换。
// 一维索引 ↔ 二维坐标的转换公式：
// 对于 m × n 的矩阵：
// 一维索引 → 二维坐标：
//   行 = 索引 / 列数
//   列 = 索引 % 列数
// 二维坐标 → 一维索引：
//   索引 = 行 × 列数 + 列
func searchMatrix74(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0]) // 行数，列数
	low, high := 0, m*n-1               // 一维数组的左右边界
	for low <= high {
		mid := low + (high-low)>>1 // 中间位置
		// 行列坐标转换
		row := mid / n
		col := mid % n
		midValue := matrix[row][col]

		if midValue == target {
			return true
		} else if midValue > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
