package leetcode

// 寻找两个正序数组的中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

// 二分搜索
// 要求算法的时间复杂度为 O(log (m+n))，不可以合并再找中位数，那样是O(m+n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	/*
			思路：
		    在两个数组上同时进行二分搜索，每次排除掉不可能包含中位数的一部分元素
			具体流程：
			1. 确保第一个数组长度 ≤ 第二个数组长度（交换以便处理）
			2. 在两个数组上定义分割线：
			   在数组A中选一个位置i
			   在数组B中选一个位置j，使得i + j = (m+n+1)/2
			   这样分割线左边的元素数量等于右边（或左边多1）
			3. 检查分割线是否满足条件：
			   需要满足：A[i-1] ≤ B[j] 且 B[j-1] ≤ A[i]
			   这意味着分割线左边的所有元素都小于等于右边的所有元素
			4. 根据比较结果调整二分搜索：
			   如果A[i-1] > B[j]：说明i太大，需要在A的左半部分继续搜索
			   如果B[j-1] > A[i]：说明i太小，需要在A的右半部分继续搜索
	*/

	// 确保第一个数组长度 <= 第二个数组长度
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	totalLeft := (m + n + 1) / 2 // 左半部分应有的元素数量（奇偶统一处理）

	// 在较短的数组nums1上使用二分搜索
	left, right := 0, m
	for left < right {
		i := left + (right-left+1)/2 // nums1的分割点
		j := totalLeft - i           // nums2的分割点

		// 检查是否分割正确
		if nums1[i-1] > nums2[j] {
			// 需要减小i，在左半部分继续搜索
			right = i - 1
		} else {
			// 需要增大i，在右半部分继续搜索
			left = i
		}
	}

	i := left
	j := totalLeft - i

	// 处理边界情况，确定四个关键值；如果越界，就用极小值/极大值替代
	var nums1LeftMax, nums1RightMin, nums2LeftMax, nums2RightMin int

	// 处理nums1的边界
	// nums1全在右边：分割点在nums1的最左边，意味着nums1没有元素在左边，此时nums1LeftMax应该是负无穷，因为左边没有元素
	if i == 0 {
		nums1LeftMax = -1 << 31 // 左边没有元素，设为负无穷math.MinInt32
	} else {
		nums1LeftMax = nums1[i-1] // 正常情况，取左边最大值
	}

	// nums1全在左边：分割点在nums1的最右边，意味着nums1没有元素在右边，此时nums1RightMin应该是正无穷，因为右边没有元素
	if i == m {
		nums1RightMin = 1<<31 - 1 // 右边没有元素，设为正无穷math.MaxInt32
	} else {
		nums1RightMin = nums1[i] // 正常情况，取右边最小值
	}

	// 处理nums2的边界
	// nums2全在右边：分割点在nums2的最左边，意味着nums2没有元素在左边，此时nums2LeftMax应该是负无穷
	if j == 0 {
		nums2LeftMax = -1 << 31 // 左边没有元素，设为负无穷math.MinInt32
	} else {
		nums2LeftMax = nums2[j-1] // 正常情况，取左边最大值
	}

	// nums2全在左边：分割点在nums2的最右边，意味着nums2没有元素在右边，此时nums2RightMin应该是正无穷
	if j == n {
		nums2RightMin = 1<<31 - 1 // 右边没有元素，设为正无穷math.MaxInt32
	} else {
		nums2RightMin = nums2[j] // 正常情况，取右边最小值
	}

	// 根据总长度奇偶性返回结果
	if (m+n)%2 == 1 {
		// 奇数长度：中位数是分割线左边最大的元素
		return float64(max(nums1LeftMax, nums2LeftMax))
	} else {
		// 偶数长度：中位数是分割线左边最大和右边最小元素的平均值
		return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2.0
	}
}
