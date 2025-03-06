package leetcode

// 加油站
// https://leetcode.cn/problems/gas-station/

// 贪心
func canCompleteCircuit(gas []int, cost []int) int {
	// 首先检查第 0 个加油站，并试图判断能否环绕一周；如果不能，就从第一个无法到达的加油站开始继续检查
	total, curSum, start := 0, 0, 0 // curSum为当前起点到当前位置的累计油量
	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		total += diff
		curSum += diff
		// 在某个点油量不够时，直接从失败点的下一个站重新开始计算，因为中间的站肯定都不行
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	// 如果所有加油站的总油量<总消耗量，那么无论从哪里出发都不可能完成旅程
	if total < 0 {
		return -1
	}
	return start
}
