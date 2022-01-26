package leetcode

// 旋转图像
// https://leetcode-cn.com/problems/rotate-image
// 原地旋转
func rotate(matrix [][]int) {
	n := len(matrix)           // 列
	for i := 0; i < n/2; i++ { // 列 [0,n/2)
		for j := 0; j < (n+1)/2; j++ { // [0, (n+1)/2)
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

// 用翻转代替旋转
func rotate1(matrix [][]int) {
	n := len(matrix)
	// 水平翻转，上面的行到下面的行
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ { // [0,n)
		for j := 0; j < i; j++ { // [0,i)
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
