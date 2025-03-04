package leetcode

// 最大连续1的个数 III
// https://leetcode.cn/problems/max-consecutive-ones-iii/

// 滑动窗口
func longestOnes(nums []int, k int) int {
	ans := 0
	zeroCount := 0

	left, right := 0, 0
	for right < len(nums) {
		// 如果当前是0，增加0的计数
		if nums[right] == 0 {
			zeroCount++
		}

		// 如果0的数量超过k，移动左指针
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// 更新最大长度
		currentLength := right - left + 1
		if currentLength > ans {
			ans = currentLength
		}

		right++
	}
	return ans
}
