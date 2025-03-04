package leetcode

import "sort"

// 摆动排序 II
// https://leetcode.cn/problems/wiggle-sort-ii/

// 排序
func wiggleSort(nums []int) {
	n := len(nums)
	arr := append([]int{}, nums...)

	// 将数组进行排序，然后从中间位置进行等分（如果数组长度为奇数，则将中间的元素分到前面），然后将两个数组进行穿插
	// 对于数组[1,5,2,4,3]，将其排序得到[1,2,3,4,5]，然后将其分割为[1, 2, 3]和[4, 5]，对两个数组进行穿插，得到[1, 4, 2, 5, 3]
	sort.Ints(arr)
	x := (n + 1) / 2 // 找到 nums 的中位数

	for i, j, k := 0, x-1, n-1; i < n; i += 2 {
		nums[i] = arr[j]
		if i+1 < n {
			nums[i+1] = arr[k]
		}
		j--
		k--
	}
}
