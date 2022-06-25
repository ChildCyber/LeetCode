package leetcode

import "sort"

// 摆动排序 II
// https://leetcode.cn/problems/wiggle-sort-ii/
// 排序
func wiggleSort(nums []int) {
	n := len(nums)
	arr := append([]int{}, nums...)
	sort.Ints(arr)

	x := (n + 1) / 2
	for i, j, k := 0, x-1, n-1; i < n; i += 2 {
		nums[i] = arr[j]
		if i+1 < n {
			nums[i+1] = arr[k]
		}
		j--
		k--
	}
}
