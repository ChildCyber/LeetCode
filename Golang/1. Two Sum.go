package leetcode

// 两数之和
// https://leetcode-cn.com/problems/two-sum/
// 暴力穷举
func twoSumForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int { // 返回和为目标值的数，数组中的下标
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
