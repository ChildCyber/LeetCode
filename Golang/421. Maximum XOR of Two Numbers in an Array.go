package leetcode

// 数组中两个数的最大异或值
// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/
// 暴力
func findMaximumXORForce(nums []int) (x int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			x = max(x, nums[i]^nums[j])
		}
	}
	return
}

// 哈希表
func findMaximumXOR(nums []int) (x int) {
	// 如果 a ^ b = c 成立，那么a ^ c = b 与 b ^ c = a 均成立
	const highBit = 30 // 最高位的二进制位编号为 30
	// 将每一个数表示为一个长度为 31 位的二进制数
	// 对于 nums[j] 而言，可以从其二进制表示中的最高位开始往低位找，尽量让每一位的异或结果为 1，这样找到的 nums[i] 与 nums[j] 的异或结果才是最大的
	for k := highBit; k >= 0; k-- {
		// 将所有的 pre^k(a_j) 放入哈希表中
		seen := map[int]bool{}
		for _, num := range nums {
			// 如果只想保留从最高位开始到第 k 个二进制位为止的部分
			// 只需将其右移 k 位
			seen[num>>k] = true
		}

		// 目前 x 包含从最高位开始到第 k+1 个二进制位为止的部分
		// 将 x 的第 k 个二进制位置为 1，即为 x = x*2+1
		xNext := x*2 + 1
		found := false

		// 枚举 i
		for _, num := range nums {
			if seen[num>>k^xNext] {
				found = true
				break
			}
		}

		if found {
			x = xNext
		} else {
			// 如果没有找到满足等式的 a_i 和 a_j，那么 x 的第 k 个二进制位只能为 0，即为 x = x*2
			x = xNext - 1
		}
	}
	return
}
