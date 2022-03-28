package leetcode

// 丢失的数字
// https://leetcode.cn/problems/missing-number/
// 异或
func missingNumber(nums []int) int {
	xor, i := 0, 0
	for i = 0; i < len(nums); i++ {
		xor = xor ^ i ^ nums[i]
	}
	return xor ^ i
}

// 排序
// 哈希表
