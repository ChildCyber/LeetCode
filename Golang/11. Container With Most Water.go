package leetcode

// 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water
// 双指针
func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1 // 首尾双指针
	for start < end {
		width := end - start // 宽
		high := 0
		// 高取左右两边最小的那个
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high // 容积
		if temp > max {
			max = temp
		}
	}
	return max
}
