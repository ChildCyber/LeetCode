package leetcode

// 旋转图像
// https://leetcode-cn.com/problems/rotate-image
// 原地旋转，将图像顺时针旋转 90 度
func rotateInPlace(matrix [][]int) {
	// 第 i 行 -> 第 n-1-i 列
	// 第 j 列 -> 第 j 行
	// 对于矩阵中的元素 matrix[i][j]，在旋转后，它的新位置为 matrix[j][n−i−1]
	n := len(matrix)
	// 分别以矩阵左上角 1/4 的各元素为起始点执行以上旋转操作，即可完整实现矩阵旋转
	for i := 0; i < n/2; i++ { // 行 [0,n/2)
		for j := 0; j < (n+1)/2; j++ { // 列 [0, (n+1)/2)
			// 首尾相接，元素旋转
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

// 用翻转代替旋转
func rotateFlip(matrix [][]int) {
	n := len(matrix)
	// 水平翻转，枚举矩阵上半部分的元素，和下半部分的元素进行交换
	for i := 0; i < n/2; i++ { // 整行替换
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转，枚举对角线左侧的元素，和右侧的元素进行交换
	for i := 0; i < n; i++ { // [0,n)
		for j := 0; j < i; j++ { // [0,i)
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
