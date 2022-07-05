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
		for num != 0 { // 计算数字每一位的平方和
			sum += (num % 10) * (num % 10)
			num = num / 10
		}

		if _, ok := record[sum]; !ok { // 该数字没出现过，添加
			if sum == 1 { // 平方和为1，该数是快乐数
				return true
			}

			record[sum] = sum
			num = sum
			sum = 0
			continue
		} else { // 该数字出现过，无限循环，但始终变不到 1
			return false
		}
	}
}

// 快慢指针
func isHappyTwo(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
