package leetcode

// 颠倒二进制位
// https://leetcode.cn/problems/reverse-bits/

// 逐位颠倒
func reverseBits(num uint32) uint32 {
	// 把num往右移动，不断的消除右边最低位的1，将这个1给ans，ans不断的左移即可
	var ans uint32
	for i := 0; i < 32; i++ {
		// (n & 1) 得到当前n的最低（最右）一位，<< (31 - i) 左移 31 - i 位，也就是把最低一位逆序添加到返回值ans对应的高位，每循环一次和之前结果进行一次或运算 |=，相当于整合了每一位的结果
		// ans |= num & 1 << (31 - i)

		// 从低位往高位枚举 num 的每一位，将其倒序添加到翻转结果 ans 中
		// 把 ans 左移一位，把 num 的二进制末尾数字，拼接到结果 ans 的末尾
		ans = ans<<1 | num&1 // num&1得到当前num的最低（最右）一位，ans<<1是ans左移一位，最低位补0，通过位或|，逐位计算，相当于将num&1追加到ans的末尾

		// 每枚举一位就将 num 右移一位，这样当前 num 的最低位就是要枚举的比特位。当 num 为 0 时即可结束循环
		num >>= 1 // num右移
	}
	return ans
}
