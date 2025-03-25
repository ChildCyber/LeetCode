package leetcode

// 两数之和
// https://leetcode-cn.com/problems/two-sum/

// 暴力
// 两层循环：外层循环遍历每一个数字作为第一个加数，内层循环遍历它之后的所有数字作为第二个加数，检查每一对数字的和是否等于target
func twoSumBrute(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 哈希表
// 用哈希表记录遍历过的数字，对于当前数字，在表中查找其补数（目标值-当前值）
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // key：数组中的数字，value：该数字在数组中的下标
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
