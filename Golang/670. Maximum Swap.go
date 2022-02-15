package leetcode

import "strconv"

// 最大交换
// https://leetcode-cn.com/problems/maximum-swap/
// 贪心算法
func maximumSwap(num int) int {
	str := strconv.Itoa(num)
	strs := []byte(str)

	// 记录每个数字出现的最后一次出现的下标
	last := [10]int{}
	for i := 0; i < len(strs); i++ {
		last[int(strs[i]-'0')] = i
	}
	// 从左向右扫描，找到当前位置右边的最大的数字并交换
	for i := 0; i < len(strs); i++ {
		// 找最大，所以倒着找
		for j := 9; j >= 0; j-- {
			if last[j] > i && j > int(strs[i]-'0') { // j在strs[i]的后面
				// 只允许交换一次，因此直接返回
				strs[last[j]], strs[i] = strs[i], strs[last[j]]
				num := string(strs)
				n, _ := strconv.Atoi(num)
				return n
			}
		}
	}
	return num
}
