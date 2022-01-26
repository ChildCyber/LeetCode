package leetcode

// 柠檬水找零
// https://leetcode-cn.com/problems/lemonade-change/
// 贪心
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 { // 5
			five++
		} else if bill == 10 { // 10
			if five == 0 {
				return false
			}
			five--
			ten++
		} else { // 20
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
