package leetcode

// 最大连续1的个数 III
// https://leetcode.cn/problems/max-consecutive-ones-iii/

// 暴力
// 思路：枚举所有子数组；数一数里面有多少个0；如果0的个数≤k，则更新答案

// 滑动窗口
// 思路：用滑动窗口维护「最多k个0」
func longestOnes(nums []int, k int) int {
	left, ans, zeroCount := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		// 如果当前是0，增加0的计数
		if nums[right] == 0 {
			zeroCount++
		}

		// 当窗口中0的数量超过k，收缩左边界
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// 更新最大合法窗口长度
		currentLength := right - left + 1 // [left, right] 左闭右闭区间
		if currentLength > ans {
			ans = currentLength
		}
	}

	return ans
}
