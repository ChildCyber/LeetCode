package leetcode

import "sort"

// 合并两个有序数组
// https://leetcode-cn.com/problems/merge-sorted-array/
// 直接合并后排序
func merge1(nums1 []int, m int, nums2 []int, _ int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 双指针
func merge(nums1 []int, m int, nums2 []int, n int) { // 非递减顺序 排列的整数数组
	sorted := make([]int, 0, m+m)
	p1, p2 := 0, 0
	for {
		if p1 == m { // nums1遍历完，nums2追加
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n { // nums2遍历完，nums1追加
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		// 都未遍历完，指针向后移动
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted) // sorted 合并后存储在数组 nums1 中
}

// 逆向双指针
func merge2(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 { // 遍历完
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}
