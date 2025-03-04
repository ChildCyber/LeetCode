package leetcode

// 爱吃香蕉的珂珂
// https://leetcode.cn/problems/koko-eating-bananas/

// 二分
func minEatingSpeed(piles []int, h int) int {
	// 找到香蕉堆中的最大值，作为右边界
	maxPile := 0
	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	// 二分查找的左右边界
	left, right := 1, maxPile

	for left < right {
		mid := left + (right-left)/2

		// 计算以当前速度 mid 吃完所有香蕉需要的时间
		time := 0
		for _, pile := range piles {
			// 向上取整的技巧： (pile + mid - 1) / mid
			time += (pile + mid - 1) / mid
		}

		if time <= h {
			// 当前速度可行，尝试更慢的速度（在左侧寻找）
			right = mid
		} else {
			// 当前速度太慢，需要更快的速度（在右侧寻找）
			left = mid + 1
		}
	}

	// 循环结束时 left == right，就是最小的可行速度
	return left
}
