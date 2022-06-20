package leetcode

// 颠倒二进制位
// https://leetcode.cn/problems/reverse-bits/
// 逐位颠倒
func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		// 把num往右移动，不断的消除右边最低位的1，将这个1给ans，ans不断的左移即可
		ans = ans<<1 | num&1 // num&1得到当前num的最低（最右）一位
		num >>= 1            // num右移
	}
	return ans
}
