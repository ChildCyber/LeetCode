package leetcode

// 搜索二维矩阵 II
// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
// Z 字形查找
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for col >= 0 && row <= len(matrix)-1 { // 从矩阵 matrix 的右上角 (0,n−1) 进行搜索
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

func searchMatrix2BS(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for _, row := range matrix {
		low, high := 0, len(matrix[0])
		for low <= high {
			mid := low + (low+high)>>1
			if row[mid] > target {
				high = mid - 1
			} else if row[mid] < target {
				low = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}
