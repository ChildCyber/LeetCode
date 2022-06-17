package leetcode

// 颠倒二进制位
// https://leetcode.cn/problems/reverse-bits/
// 逐位颠倒
func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans = ans<<1 | num&1
		num >>= 1
	}
	return ans
}
