package leetcode

import "sort"

// 丢失的数字
// https://leetcode.cn/problems/missing-number/
// 只有一个数字丢失，返回丢失的数字

// 异或
func missingNumber(nums []int) int {
	xor, i := 0, 0
	for i = 0; i < len(nums); i++ {
		xor = xor ^ i ^ nums[i] // 除了丢失的数，每个数都出现2次
	}
	return xor ^ i
}

// 排序
func missingNumberSort(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if num != i {
			return i
		}
	}
	return len(nums)
}

// 哈希表
func missingNumberMap(nums []int) int {
	has := map[int]bool{}
	for _, v := range nums {
		has[v] = true
	}
	for i := 0; ; i++ {
		if !has[i] {
			return i
		}
	}
}
