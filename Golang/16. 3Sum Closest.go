package leetcode

import "sort"

// 最接近的三数之和
// https://leetcode.cn/problems/3sum-closest/

// 排序+双指针
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2] // 初始值
	length := len(nums)

	for i := 0; i < length-2; i++ {
		// 跳过重复元素（可选，但可以优化性能）
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, length-1
		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			// 如果正好等于target，直接返回
			if currentSum == target {
				return target
			}

			// 更新最接近的和
			if abs(currentSum-target) < abs(closestSum-target) {
				closestSum = currentSum
			}

			// 移动指针
			if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closestSum
}
