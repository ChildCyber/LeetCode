package leetcode

// 或运算的最小翻转次数
// https://leetcode.cn/problems/minimum-flips-to-make-a-or-b-equal-to-c/

func minFlips(a int, b int, c int) int {
	flips := 0
	// 逐位处理，直到三个数都为 0
	for a > 0 || b > 0 || c > 0 {
		// 取出当前正在处理的最低位（二进制最后一位）
		ai := a & 1
		bi := b & 1
		ci := c & 1

		if ci == 1 {
			// 目标是 ai | bi == 1
			// 若 ai==0 且 bi==0，需要一次翻转（把 a 或 b 的该位变成 1）
			if ai == 0 && bi == 0 {
				flips++
			}
		} else {
			// ci == 0, 目标是 ai == 0 且 bi == 0
			// 若 ai==1 && bi==1 -> 需要两次翻转
			// 若 ai==1 xor bi==1 -> 需要一次翻转
			if ai == 1 && bi == 1 {
				flips += 2
			} else if ai == 1 || bi == 1 {
				flips++
			}
		}

		// 处理下一位
		a >>= 1
		b >>= 1
		c >>= 1
	}
	return flips
}
