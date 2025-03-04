package leetcode

// 最近的请求次数
// https://leetcode.cn/problems/number-of-recent-calls/

// 队列
type RecentCounter struct {
	requests []int
}

func Constructor933() RecentCounter {
	return RecentCounter{
		requests: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	// 将当前请求时间加入队列
	this.requests = append(this.requests, t)

	// 移除所有不在时间窗口 [t-3000, t] 内的请求
	for len(this.requests) > 0 && this.requests[0] < t-3000 {
		this.requests = this.requests[1:]
	}

	// 返回当前时间窗口内的请求数量
	return len(this.requests)
}
