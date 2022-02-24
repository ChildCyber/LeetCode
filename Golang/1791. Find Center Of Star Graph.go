package leetcode

// 找出星型图的中心节点
// https://leetcode-cn.com/problems/find-center-of-star-graph/
// 计算每个节点的度
func findCenter(edges [][]int) int {
	// 遍历 edges 中的每条边并计算每个节点的度，度为 n−1 的节点即为中心节点
	n := len(edges) + 1
	degrees := make([]int, n+1)
	for _, e := range edges {
		degrees[e[0]]++
		degrees[e[1]]++
	}
	for i, d := range degrees {
		if d == n-1 {
			return i
		}
	}
	return -1
}
