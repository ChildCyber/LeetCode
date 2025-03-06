package leetcode

// 删掉一个元素以后全为 1 的最长子数组
// https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/

// 暴力
// 枚举要删除的每个下标；删除它，计算最长连续1的长度；取最大值

// 滑动窗口
// LC1493 = LC1004(k=1) 且 必须操作一次
func longestSubarray(nums []int) int {
	// 删除一个0 ≈ 允许窗口内最多有一个0
	left, right := 0, 0
	zeroCount := 0
	ans := 0

	// 只要窗口内的zeroCount <= 1，窗口就有效
	for right < len(nums) {
		// 如果当前是0，增加0的计数
		if nums[right] == 0 {
			zeroCount++
		}

		// 如果0的数量超过1，移动左指针
		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// 更新最大长度
		// 注意：这里计算的是删除一个元素后的长度，所以是(right-left)而不是(right-left+1)
		currentLength := right - left
		if currentLength > ans {
			ans = currentLength
		}

		right++
	}

	return ans
}

func longestSubarray1(nums []int) int {
	left, zeroCount, ans := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		// 窗口无效（0 超过 1 个），收缩左边界
		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// 更新最大窗口长度
		ans = max(ans, right-left+1)
	}

	// 题目要求“必须删掉一个元素”，哪怕都是1
	return ans - 1
}
