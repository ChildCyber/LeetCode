package leetcode

// 存在重复元素 II
// https://leetcode.cn/problems/contains-duplicate-ii/

// 暴力
func containsNearbyDuplicateBrute(nums []int, k int) bool {
	// 对每个元素，向后看 k 个位置，看是否有重复
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// 滑动窗口+哈希表
func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int) // 存储元素和最新索引

	for i, num := range nums {
		// 如果元素存在且索引差 <= k
		if prevIndex, ok := indexMap[num]; ok && i-prevIndex <= k {
			return true
		}
		// 更新元素的最新索引
		indexMap[num] = i
	}

	return false
}
