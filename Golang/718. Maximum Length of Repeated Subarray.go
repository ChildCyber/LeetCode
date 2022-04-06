package leetcode

// 最长重复子数组
// https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/
// 动态规划
func findLengthDP(A []int, B []int) int {
	n, m := len(A), len(B)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

// 滑动窗口
func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	ret := 0
	for i := 0; i < n; i++ {
		length := min(m, n-i)
		maxLen := maxLength(A, B, i, 0, length)
		ret = max(ret, maxLen)
	}
	for i := 0; i < m; i++ {
		length := min(n, m-i)
		maxLen := maxLength(A, B, 0, i, length)
		ret = max(ret, maxLen)
	}
	return ret
}

func maxLength(A, B []int, addA, addB, len int) int {
	ans, k := 0, 0
	for i := 0; i < len; i++ {
		if A[addA+i] == B[addB+i] {
			k++
		} else {
			k = 0
		}
		ans = max(ans, k)
	}
	return ans
}
