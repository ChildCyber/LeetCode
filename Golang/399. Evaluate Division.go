package leetcode

// 除法求值
// https://leetcode.cn/problems/evaluate-division/

// 问题本质：带权有向图的路径查询
// 图的抽象：
//   顶点：每一个独特的变量（如 'a', 'b', 'c' 等）
//   边：把每个等式当作一条有向边。关系 A/B=2.0 意味着从 A 到 B 有一条有向边
// 问题建模：
//   查询相当于在图中找一条从start到end的路径，路径边权连乘即结果

// 图遍历DFS
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 用邻接表表示有向带权图： key: 节点名（变量名），value: map[邻居]权重（cur -> nei 的值）
	graph := make(map[string]map[string]float64)

	// 1. 构建图
	// 对每个等式 a/b=val，加入两条边：
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

	// 2. 定义DFS
	var dfs func(start, end string, visited map[string]bool, acc float64) float64
	// acc: 当前累积乘积（从查询的起点到 cur 的乘积）
	dfs = func(start, end string, visited map[string]bool, acc float64) float64 {
		// 如果当前节点在图中都不存在，则说明没法出发（孤立节点）
		if _, ok := graph[start]; !ok {
			return -1.0
		}
		// 找到目标，直接返回当前累积值（注意：如果 start==end 并且节点存在，应返回 1.0）
		if start == end {
			return acc
		}

		// 标记已访问，避免环路导致无限递归
		visited[start] = true
		// 遍历所有邻居：start -> nei 的权重为 val
		for nei, val := range graph[start] {
			if !visited[nei] {
				// 递归：将当前累积乘以该边权
				res := dfs(nei, end, visited, acc*val)
				// 如果子路径找到了（res != -1.0），直接向上返回最终结果（短路）
				if res != -1.0 {
					return res
				}
			}
		}
		return -1.0
	}

	// 3. 处理queries
	ans := make([]float64, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		ans[i] = dfs(start, end, map[string]bool{}, 1.0)
	}
	return ans
}
