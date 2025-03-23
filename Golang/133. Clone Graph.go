package leetcode

// 克隆图
// https://leetcode.cn/problems/clone-graph/

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

// DFS+哈希表
func cloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}

	// 哈希表：原节点 -> 克隆节点
	visited := make(map[*GraphNode]*GraphNode)

	// 定义DFS闭包函数
	var dfs func(*GraphNode) *GraphNode
	dfs = func(n *GraphNode) *GraphNode {
		// 如果节点已经访问过，直接返回对应的克隆节点
		if cloned, exists := visited[n]; exists {
			return cloned
		}

		// 创建当前节点的克隆
		cloneNode := &GraphNode{Val: n.Val, Neighbors: make([]*GraphNode, 0)}
		visited[n] = cloneNode

		// 递归克隆所有邻居节点
		for _, neighbor := range n.Neighbors {
			clonedNeighbor := dfs(neighbor)
			cloneNode.Neighbors = append(cloneNode.Neighbors, clonedNeighbor)
		}

		return cloneNode
	}

	return dfs(node)
}

// BFS+哈希表
func cloneGraphBFS(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}

	// 哈希表：原节点 -> 克隆节点
	visited := make(map[*GraphNode]*GraphNode)

	// 创建根节点的克隆
	cloneNode := &GraphNode{Val: node.Val, Neighbors: []*GraphNode{}}
	visited[node] = cloneNode

	// BFS队列
	queue := []*GraphNode{node}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// 遍历当前节点的所有邻居
		for _, neighbor := range curr.Neighbors {
			// 如果邻居节点还没有被克隆
			if _, exists := visited[neighbor]; !exists {
				// 创建邻居的克隆
				cloneNeighbor := &GraphNode{Val: neighbor.Val, Neighbors: []*GraphNode{}}
				visited[neighbor] = cloneNeighbor
				queue = append(queue, neighbor)
			}
			// 将克隆的邻居添加到当前克隆节点的邻居列表中
			visited[curr].Neighbors = append(visited[curr].Neighbors, visited[neighbor])
		}
	}

	return cloneNode
}
