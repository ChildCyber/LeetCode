package leetcode

import "sort"

// 三数之和
// https://leetcode-cn.com/problems/3sum/
// 暴力
// 三重循环，分别找出三个元素再判断和是否为0

// 排序 + 双指针
func threeSum(nums []int) [][]int {
	// 排序，减少重复
	sort.Ints(nums)
	// 双指针
	// 第二重循环和第三重循环实际上是并列的关系
	// 保持第二重循环不变，而将第三重循环变成一个从数组最右端开始向左移动的指针
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	// 第一重循环
	for index = 1; index < length-1; index++ { // 注意：从1开始
		start, end = 0, length-1
		// 只有和上一次枚举的元素不相同，才会进行枚举
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}

		for start < index && end > index {
			// 第三重循环「跳到」下一个不相同的元素
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}

			// 相加和
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 { // 右指针左移
				end--
			} else { // 左指针右移
				start++
			}
		}
	}
	return result
}

func threeSum1(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	sort.Ints(nums) // 排序

	// 枚举 a
	for first := 0; first < n; first++ { // 注意：从0开始
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] { // 为了防止数组越界，需要跳过下标为0的情况
			continue
		}

		// 确定了第一个元素a，剩下的就是two sum
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first] // 三数相加和为0
		// 枚举 b
		for second := first + 1; second < n; second++ { // 从first右边开始
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
