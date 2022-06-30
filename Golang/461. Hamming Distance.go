package leetcode

// 汉明距离
// https://leetcode-cn.com/problems/hamming-distance/
// 两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目
// 位计数问题

// 移位实现位计数
func hammingDistance(x int, y int) int {
	distance := 0
	for xor := x ^ y; xor != 0; xor &= (xor - 1) { // 利用异或得到二进制位不同的数，再利用 x &= (x-1) 不断的清除最低位的 1
		distance++
	}
	return distance
}

// 内置位计数功能
// return bits.OnesCount(uint(x ^ y))
