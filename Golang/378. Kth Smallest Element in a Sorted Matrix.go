package leetcode

import "sort"

// 有序矩阵中第 K 小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/

// 直接排序
func kthSmallestSort(matrix [][]int, k int) int {
	// 将二维数组转成一维数组
	rows, columns := len(matrix), len(matrix[0])
	sorted := make([]int, rows*columns)
	index := 0
	for _, row := range matrix {
		for _, num := range row {
			sorted[index] = num
			index++
		}
	}
	// 对该一维数组进行排序
	sort.Ints(sorted)
	// 最后这个一维数组中的第 k 个数即为答案
	return sorted[k-1]
}

// 二分查找
func kthSmallest378(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if check(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(matrix [][]int, mid, k, n int) bool {
	i, j := n-1, 0
	num := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			num += i + 1
			j++
		} else {
			i--
		}
	}
	return num >= k
}
