package leetcode

// 猜数字大小
// https://leetcode.com/problems/guess-number-higher-or-lower/

// 二分
func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2 // 不是 (left + right) / 2，避免大数相加溢出
		result := guess(mid)

		if result == 0 {
			return mid
		} else if result == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1 // 理论上不会执行到这里
}
