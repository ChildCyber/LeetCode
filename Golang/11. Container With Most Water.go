package leetcode

// 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water
// 暴力：遍历所有可能的左右板子组合，计算面积，记录最大值。1使用两个循环，外层循环选择左边的板子。2内层循环选择右边的板子（从左边板子的右侧开始）。3.对于每对板子(i, j)：计算宽度：j - i；计算高度：min(height[i], height[j])；计算面积：宽度 × 高度4.记录所有面积中的最大值
// 时间复杂度：O(n^2)

// 双指针
func maxArea(height []int) int {
	// 容器的容量由两个因素决定：宽度和高度（较短板子的高度）
	// 向内移动指针时，宽度一定会减小，要增加容量，唯一的可能性就是增加高度
	ans, start, end := 0, 0, len(height)-1 // 首尾双指针
	for start < end {
		width := end - start // 宽
		high := 0
		// 高度取左右两边较小的，同时移动较矮的指针
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high // 容积
		if temp > ans {
			ans = temp
		}
	}
	return ans
}
