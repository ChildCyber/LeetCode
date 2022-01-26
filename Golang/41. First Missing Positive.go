package leetcode

// 缺失的第一个正数
// https://leetcode-cn.com/problems/first-missing-positive
func firstMissingPositive(nums []int) int {
	numMap := make(map[int]int, len(nums)) // 无额外空间复杂度要求
	// 思路：将数组所有的数放入哈希表，随后从 1 开始依次枚举正整数，并判断其是否在哈希表中；
	for _, v := range nums { // 将数组所有的数放入map
		numMap[v] = v
	}
	// 最小的正整数从1开始
	// i为nums中的index，根据i+1的key到map中查询
	for i := 0; i < len(nums); i++ { // 从1开始，依次对比map中是否存在i
		if _, ok := numMap[i+1]; !ok {
			return i + 1
		}
	}
	return len(nums) + 1
}
