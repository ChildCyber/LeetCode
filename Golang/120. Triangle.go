package leetcode

import "math"

// 三角形最小路径和
// https://leetcode-cn.com/problems/triangle/
// 动态规划
func minimumTotal(triangle [][]int) int {
	// 用 f[i][j] 表示从三角形顶部走到位置 (i,j) 的最小路径和。位置 (i,j) 指的是三角形中第 i 行第 j 列（均从 0 开始编号）的位置。
	// 状态转移方程为：f[i][j]=min(f[i−1][j−1],f[i−1][j])+c[i][j]
	// 其中 c[i][j] 表示位置 (i,j) 对应的元素值。
	n := len(triangle)
	f := make([][]int, n) // n*n，三角形的行数
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i][0] = f[i-1][0] + triangle[i][0] // 第 i 行的最左侧时，只能从第 i−1 行的最左侧移动过来
		for j := 1; j < i; j++ {
			f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
		}
		f[i][i] = f[i-1][i-1] + triangle[i][i] // 第 i 行的最右侧时，只能从第 i−1 行的最右侧移动过来
	}

	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = min(ans, f[n-1][i])
	}
	return ans
}
