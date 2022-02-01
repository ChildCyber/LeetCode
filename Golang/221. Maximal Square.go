package leetcode

// 最大正方形
// https://leetcode-cn.com/problems/maximal-square/
// 暴力, 找到最大正方形的边长
func maximalSquare(matrix [][]byte) int {
	maxSide := 0 // 最大边长
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxSide
	}

	rows, columns := len(matrix), len(matrix[0]) // 矩阵长，宽
	// 遍历矩阵中的每个元素，每次遇到 1，则将该元素作为正方形的左上角
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' { // 遇到一个 1，作为正方形的左上角
				// 计算可能的最大正方形边长
				maxSide = max(maxSide, 1)
				curMaxSide := min(rows-i, columns-j) // 根据左上角所在的行和列计算可能的最大正方形的边长，matrix[i][j]
				for k := 1; k < curMaxSide; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' { // 正方形右下角
						break
					}
					// 判断新增的一行一列是否均为 1
					for m := 0; m < k; m++ {
						if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSide = max(maxSide, k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSide * maxSide // 最大面积
}

// 动态规划
func maximalSquareDP(matrix [][]byte) int {
	// 用 dp(i,j) 表示以 (i,j) 为右下角，且只包含 1 的正方形的边长最大值
	dp := make([][]int, len(matrix))
	maxSide := 0 // 最大边长
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ { // 从1开始
		for j := 1; j < len(matrix[i]); j++ { // 从1开始
			if dp[i][j] == 1 { // 该位置的值是 1，则 dp(i,j) 的值由其上方、左方和左上方的三个相邻位置的 dp 值决定
				// 当前位置的元素值等于三个相邻位置的元素中的最小值加 1
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}
