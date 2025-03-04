package leetcode

// 重新规划路线
// https://leetcode.cn/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/

// DFS
func minReorder(n int, connections [][]int) int {
	// 建图：用邻接表存 (创建邻接表表示城市之间的连接关系，同时记录边的方向信息)
	graph := make([][][2]int, n)
	for _, conn := range connections {
		from, to := conn[0], conn[1]
		graph[from] = append(graph[from], [2]int{to, 1}) // 正向：from→to，需要反转
		graph[to] = append(graph[to], [2]int{from, 0})   // 反向：to→from，不需反转
	}

	visited := make([]bool, n)
	count := 0

	var dfs func(cur int)
	dfs = func(cur int) {
		visited[cur] = true
		for _, next := range graph[cur] {
			nextNode, needsReverse := next[0], next[1]
			if !visited[nextNode] { // 邻居未被访问且是正向边
				// 如果方向是正向(1)，说明需要改变这条边
				// 因为实际道路是从当前节点指向邻居，但需要的是从邻居指向当前节点
				count += needsReverse
				dfs(nextNode)
			}
		}
	}

	dfs(0)
	return count
}
