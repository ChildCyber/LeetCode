package leetcode

// 接雨水
// https://leetcode-cn.com/problems/trapping-rain-water/
// 双指针
func trap(height []int) int {
	// 维护两个指针 left 和 right，以及两个变量 leftMax 和 rightMax，初始时 left=0,right=n−1,leftMax=0,rightMax=0。
	// 指针 left 只会向右移动，指针 right 只会向左移动，在移动指针的过程中维护两个变量 leftMax 和 rightMax 的值。
	// 当两个指针没有相遇时，进行如下操作：
	// 使用 height[left] 和 height[right] 的值更新 leftMax 和 rightMax 的值；
	// 如果 height[left]<height[right]，则必有 leftMax<rightMax，下标 left 处能接的雨水量等于 leftMax−height[left]，将下标 left 处能接的雨水量加到能接的雨水总量，然后将 left 加 1（即向右移动一位）；
	// 如果 height[left]≥height[right]，则必有 leftMax≥rightMax，下标 right 处能接的雨水量等于 rightMax−height[right]，将下标 right 处能接的雨水量加到能接的雨水总量，然后将 right 减 1（即向左移动一位）。
	// 当两个指针相遇时，即可得到能接的雨水总量。
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right { // 两指针相遇
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++ // 向右移动一位
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right-- // 向左移动一位
		}
	}
	return res
}

// 动态规划
// 单调栈
