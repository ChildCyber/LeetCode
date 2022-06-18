package leetcode

// 快乐数
// https://leetcode.cn/problems/happy-number/
// 用哈希集合检测循环
func isHappy(n int) bool {
	if n == 0 {
		return false
	}

	ans := 0
	num := n
	record := map[int]int{}
	for {
		for num != 0 {
			ans += (num % 10) * (num % 10)
			num = num / 10
		}
		if _, ok := record[ans]; !ok {
			if ans == 1 {
				return true
			}
			record[ans] = ans
			num = ans
			ans = 0
			continue
		} else {
			return false
		}
	}
}
