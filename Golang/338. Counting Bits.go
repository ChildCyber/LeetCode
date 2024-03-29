package leetcode

import "math/bits"

// 比特位计数
// https://leetcode-cn.com/problems/counting-bits/
// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

// 内置函数
func countBitsBuiltin(num int) []int {
	ans := make([]int, num+1)
	for i := 0; i <= num; i++ {
		ans[i] = bits.OnesCount(uint(i))
	}
	return ans
}

// 动态规划，最低设置位
func countBits(num int) []int {
	ans := make([]int, num+1) // dp[0] = 0，0的比特位为0
	// 遍历从 1 到 num 的每个正整数 i，计算 ans 的值。最终得到的数组 ans 即为答案
	for i := 1; i <= num; i++ { // 由于是从0-num是递增序列，则每次使用 n &= n - 1去除一位1后，必定小于原数，则必能在以前序列中找到此数，将其原有的值加上去除的1可得当前数1的个数
		ans[i] += ans[i&(i-1)] + 1
	}
	return ans
}
