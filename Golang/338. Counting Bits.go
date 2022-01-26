package leetcode

// 比特位计数
// https://leetcode-cn.com/problems/counting-bits/
// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
// Golang bits.OnesCount计算给定整数的二进制表示中的1的数目
func countBits(num int) []int { // dp，最低设置位
	bits := make([]int, num+1)  // dp[0] = 0，0的比特位为0
	for i := 1; i <= num; i++ { // 从1开始
		bits[i] += bits[i&(i-1)] + 1
	}
	return bits
}
