package leetcode

// 快乐数
// https://leetcode.cn/problems/happy-number/
// 用哈希集合检测循环
func isHappy(n int) bool {
	if n == 0 {
		return false
	}

	sum := 0
	num := n
	record := map[int]int{}
	for {
		for num != 0 {
			sum += (num % 10) * (num % 10)
			num = num / 10
		}

		if _, ok := record[sum]; !ok {
			if sum == 1 {
				return true
			}

			record[sum] = sum
			num = sum
			sum = 0
			continue
		} else {
			return false
		}
	}
}
