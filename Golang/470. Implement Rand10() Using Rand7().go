package leetcode

import "math/rand"

// 用 Rand7() 实现 Rand10()
// https://leetcode-cn.com/problems/implement-rand10-using-rand7
func rand7() int {
	return rand.Intn(7)
}

func rand10() int {
	for {
		row := rand7()
		col := rand7()
		idx := (row-1)*7 + col
		if idx <= 40 { // 拒绝采样
			return 1 + (idx-1)%10
		}
	}
}
