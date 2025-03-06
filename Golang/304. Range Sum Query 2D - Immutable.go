package leetcode

// 二维区域和检索 - 矩阵不可变
// https://leetcode.cn/problems/range-sum-query-2d-immutable/

// 暴力
// 直接遍历子矩阵相加
/*
	sum := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += matrix[i][j]
		}
	}
*/

// 前缀和-2D矩阵
type NumMatrix struct {
	preSum [][]int // preSum[i][j]表示从 (0,0) 到 (i-1,j-1) 的矩形和
}

func Constructor304(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	m, n := len(matrix), len(matrix[0])
	preSum := make([][]int, m+1) // preSum的尺寸是(m+1) × (n+1)
	for i := 0; i <= m; i++ {
		preSum[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 二维前缀和的递推公式
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	return NumMatrix{preSum}
}

func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	ps := nm.preSum
	// 区域和的计算公式
	return ps[row2+1][col2+1] - ps[row1][col2+1] - ps[row2+1][col1] + ps[row1][col1] // “+1”是因为preSum比matrix多一圈
}
