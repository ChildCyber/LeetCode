package leetcode

// 下一个更大元素 I
// https://leetcode.cn/problems/next-greater-element-i/
// 暴力
// 时间复杂度：O(mn)，其中m是nums1的长度，n是nums2的长度
// 空间复杂度：O(1)
func nextGreaterElement496Force(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	ans := make([]int, m)
	for i, num := range nums1 {
		j := 0
		for j < n && nums2[j] != num { // 直到nums1[i] == nums2[j]
			j++
		}
		k := j + 1
		for k < n && nums2[k] < nums2[j] { // 直到nums2[k] > nums2[j]
			k++
		}
		if k < n {
			ans[i] = nums2[k]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

// 单调栈 + 哈希表
func nextGreaterElement496(nums1, nums2 []int) []int {
	mp := map[int]int{}
	stack := []int{} // 用单调栈中维护当前位置右边的更大的元素列表，从栈底到栈顶的元素是单调递减的
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		// 构建单调栈
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		// 利用map记录元素值与其右边第一个更大的元素值的对应关系
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}

	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		ans[i] = mp[num]
	}

	return ans
}
