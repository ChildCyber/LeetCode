package leetcode

// 找出两数组的不同
// https://leetcode.cn/problems/find-the-difference-of-two-arrays/

// 哈希表
func findDifference(nums1 []int, nums2 []int) [][]int {
	// 创建两个集合
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	// 填充集合
	for _, num := range nums1 {
		set1[num] = true
	}
	for _, num := range nums2 {
		set2[num] = true
	}

	// 计算差集
	result1 := make([]int, 0)
	result2 := make([]int, 0)

	// nums1 - nums2
	for num := range set1 {
		if !set2[num] {
			result1 = append(result1, num)
		}
	}

	// nums2 - nums1
	for num := range set2 {
		if !set1[num] {
			result2 = append(result2, num)
		}
	}

	return [][]int{result1, result2}
}
