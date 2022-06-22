package leetcode

// hard
// 缺失的第一个正数
// https://leetcode-cn.com/problems/first-missing-positive
// 未排序的整数数组nums，找出其中没有出现的最小的正整数

// 将数组所有的数放入哈希表，随后从 1 开始依次枚举正整数，并判断其是否在哈希表中
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func firstMissingPositiveMap(nums []int) int {
	n := len(nums)
	numMap := make(map[int]int, n)
	for _, v := range nums { // 将数组中的所有的数放入map
		numMap[v] = v
	}
	// 最小的正整数从1开始，根据i+1的key到map中查询
	for i := 0; i < n; i++ { // 从1开始，依次对比map中是否存在i
		if _, ok := numMap[i+1]; !ok {
			return i + 1
		}
	}
	return n + 1
}

// 从 1 开始依次枚举正整数，并遍历数组，判断其是否在数组中

// 置换
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
