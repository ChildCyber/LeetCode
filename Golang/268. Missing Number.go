package leetcode

import "sort"

// 丢失的数字
// https://leetcode.cn/problems/missing-number/
// 在包含[0, n]共n个数的数组nums只有一个数字丢失，返回丢失的数字

// 异或
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func missingNumber(nums []int) int {
	xor, i := 0, 0
	for i = 0; i < len(nums); i++ {
		xor = xor ^ i ^ nums[i] // 除了丢失的数，每个数都出现2次，i^nums[i]==0
	}
	return xor ^ i
}

// 数学
// 高斯公式
func missingNumberMath(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}
	return total - arrSum
}

// 排序
// 时间复杂度：O(n log n)
// 空间复杂度：O(log n)
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
// 时间复杂度：O(n)
// 空间复杂度：O(n)
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
