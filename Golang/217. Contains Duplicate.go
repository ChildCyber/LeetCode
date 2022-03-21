package leetcode

// 存在重复元素
// https://leetcode-cn.com/problems/contains-duplicate/
// 哈希表
func containsDuplicate(nums []int) bool {
	record := make(map[int]bool, len(nums))
	for _, n := range nums {
		if _, found := record[n]; found {
			return true
		}
		record[n] = true
	}
	return false
}
