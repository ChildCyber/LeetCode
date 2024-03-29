package leetcode

// 位1的个数
// https://leetcode.cn/problems/number-of-1-bits/
// 汉明重量

// 位与
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
