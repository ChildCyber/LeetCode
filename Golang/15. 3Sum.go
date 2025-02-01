package leetcode

import "sort"

// 三数之和
// https://leetcode-cn.com/problems/3sum/
// 暴力
// 三重循环，分别找出三个元素再判断和是否为0
func threeSumBrute(nums []int) [][]int {
	sort.Ints(nums) // 排序，要求不可以包含重复的三元组
	ans := make([][]int, 0)
	n := len(nums)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for third := second + 1; third < n; third++ {
				if nums[first]+nums[second]+nums[third] == 0 {
					ans = append(ans, []int{nums[first], nums[second], nums[third]})
				}
			}
		}
	}
	return ans
}

// 排序+双指针
// 时间复杂度：O(N^2)，其中 N 是数组 nums 的长度
// 空间复杂度：O(logN)
func threeSum(nums []int) [][]int {
	/**
	步骤：
	1. 排序
	重复的数字会挨在一起，便于跳过重复；可以按顺序处理，不会漏掉也不会重复
	2. 固定一个数（最外层循环），把三数问题转化为两数问题
	遍历数组，把每个数当作"第一个数"，然后在它后面的数字中找另外两个数
	3. 在剩余部分使用双指针找另外两个数
	双指针技巧：
	左指针从剩余部分开头开始
	右指针从剩余部分末尾开始
	根据当前和与目标值的比较移动指针
	具体移动规则：
	当前和 < 目标值：左指针右移（增加和）
	当前和 > 目标值：右指针左移（减少和）
	当前和 = 目标值：找到一组解，同时移动两个指针
	4. 跳过重复（避免重复组合）
	跳过时机：
	固定第一个数时：如果当前数和前一个数相同，跳过
	找到一组解后：跳过左指针和右指针指向的重复数字
	**/

	// 排序，减少重复
	sort.Ints(nums)

	ans := make([][]int, 0)
	length := len(nums)
	// 外层循环，固定第一个数
	for i := 0; i < length-2; i++ { // 第一个数至少需要留出两个位置给另外两个数，如果 i 达到 length-2，那么它后面只剩下一个位置（length-1），无法形成三个数的组合
		// 跳过重复的固定数（避免重复三元组）
		if i > 0 && nums[i] == nums[i-1] { // 跳过i==0情况
			continue
		}

		left, right := i+1, length-1 // 初始化双指针
		target := -nums[i]           // 需要找到两数之和等于-target
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				// 找到一组解
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 移动指针并跳过重复
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				// 和太小，左指针右移
				left++
				// 跳过重复的左指针元素（可选优化）
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else {
				// 和太大，右指针左移
				right--
				// 跳过重复的右指针元素（可选优化）
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return ans
}
