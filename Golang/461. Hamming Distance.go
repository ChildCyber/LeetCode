package leetcode

// 汉明距离
// https://leetcode-cn.com/problems/hamming-distance/
// 两个整数之间的汉明距离是对应二进制位置上数字不同的位数
func hammingDistance(x int, y int) int {
	distance := 0
	for xor := x ^ y; xor != 0; xor &= (xor - 1) {
		distance++
	}
	return distance
}
// return bits.OnesCount(uint(x ^ y))
