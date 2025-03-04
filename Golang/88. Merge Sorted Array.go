package leetcode

import "sort"

// 合并两个有序数组
// https://leetcode-cn.com/problems/merge-sorted-array/

// 直接合并后排序
// 时间复杂度：O((m+n)log(m+n))
// 空间复杂度：O(log(m+n))
func mergeCopy(nums1 []int, m int, nums2 []int, _ int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 双指针
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n)
func merge1(nums1 []int, m int, nums2 []int, n int) {
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
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- { // 从后向前遍历，每次取两者之中的较大者放进nums1的最后面
		var cur int
		if p1 == -1 { // nums2遍历完
			cur = nums2[p2]
			p2--
		} else if p2 == -1 { // nums2遍历完
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] { // 未遍历完，比较大小
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		// 直接对数组nums1原地修改，后半部分是空的，可以直接覆盖而不会影响结果
		nums1[tail] = cur
	}
}
