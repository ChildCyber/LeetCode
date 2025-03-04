package leetcode

// 除法求值
// https://leetcode.cn/problems/evaluate-division/

// 图的DFS
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 用邻接表表示有向带权图： key: 节点名（变量名），value: map[邻居]权重（cur -> nei 的值）
	graph := make(map[string]map[string]float64)

	// 1. 建图：对每个等式 a / b = val，加入两条边：
	//	  a -> b 权重 val
	//	  b -> a 权重 1/val （便于从任意方向查询）
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		val := values[i]

		// 确保两个节点的邻接 map 存在
		if _, ok := graph[a]; !ok {
			graph[a] = make(map[string]float64)
		}
		if _, ok := graph[b]; !ok {
			graph[b] = make(map[string]float64)
		}
		// 正向边和反向边
		graph[a][b] = val
		graph[b][a] = 1.0 / val
	}

	// 2. 定义 DFS
	var dfs func(cur, target string, visited map[string]bool, acc float64) float64
	// acc: 当前累积乘积（从查询的起点到 cur 的乘积）
	dfs = func(cur, target string, visited map[string]bool, acc float64) float64 {
		// 如果当前节点在图中都不存在，则说明没法出发（孤立节点）
		if _, ok := graph[cur]; !ok {
			return -1.0
		}
		// 找到目标，直接返回当前累积值（注意：如果 start==end 并且节点存在，应返回 1.0）
		if cur == target {
			return acc
		}
		// 标记已访问，避免环路导致无限递归
		visited[cur] = true
		// 遍历所有邻居：cur -> nei 的权重为 val
		for nei, val := range graph[cur] {
			if !visited[nei] {
				// 递归：将当前累积乘以该边权
				res := dfs(nei, target, visited, acc*val)
				// 如果子路径找到了（res != -1.0），直接向上返回最终结果（短路）
				if res != -1.0 {
					return res
				}
			}
		}
		return -1.0
	}

	// 3. 处理 queries
	ans := make([]float64, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		ans[i] = dfs(start, end, map[string]bool{}, 1.0)
	}
	return ans
}
