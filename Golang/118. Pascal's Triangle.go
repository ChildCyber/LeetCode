package leetcode

// 杨辉三角
// https://leetcode.cn/problems/pascals-triangle/

// 递推
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		// 每行有i+1个元素
		ans[i] = make([]int, i+1)
		// 首尾元素为1
		ans[i][0] = 1
		ans[i][i] = 1

		// 计算中间元素
		for j := 1; j < i; j++ {
			// 递推关系：中间元素等于上一行同位置与前一位置元素之和
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
