package leetcode

// 快乐数
// https://leetcode.cn/problems/happy-number/
// 关键是检测循环

// 哈希表
func isHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {
		seen[n] = true
		n = getSquareSum(n)
	}
	return n == 1
}

// 快慢指针
func isHappyTwo(n int) bool {
	slow, fast := n, getSquareSum(n)
	for fast != 1 && slow != fast {
		slow = getSquareSum(slow)               // 算一次平方和
		fast = getSquareSum(getSquareSum(fast)) // 算两次平方和
	}
	return fast == 1
}

// 辅助函数：计算数字各位的平方和
func getSquareSum(num int) int {
	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	return sum
}
