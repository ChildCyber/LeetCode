package leetcode

// 删掉一个元素以后全为 1 的最长子数组
// https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/

// 滑动窗口
// 1004 k==1
func longestSubarray(nums []int) int {
	left, right := 0, 0
	zeroCount := 0
	ans := 0

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
		// 注意：这里计算的是删除一个元素后的长度，所以是(right - left)而不是(right - left + 1)
		currentLength := right - left
		if currentLength > ans {
			ans = currentLength
		}

		right++
	}

	return ans
}
