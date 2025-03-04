package leetcode

// 找到最高海拔
// https://leetcode.cn/problems/find-the-highest-altitude/

func largestAltitude(gain []int) int {
	ans := 0
	height := 0
	for i := 0; i < len(gain); i++ {
		height += gain[i]
		ans = max(ans, height)
	}
	return ans
}
