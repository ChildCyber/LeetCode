package leetcode

// 省份数量
// https://leetcode.cn/problems/number-of-provinces/

// 问题本质：isConnected直接就是图的邻接矩阵，统计无向图中的连通分量的数量
// 图的抽象：
//   节点：城市 0, 1, 2
//   边：isConnected[i][j]=1表示城市i和j直接相连
// 连通分量：直接或间接相连的城市集合（省份）

// 递归DFS
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n) // 记录城市是否被访问过
	count := 0

	// 从当前城市出发，标记所有连通的城市
	var dfs func(int)
	dfs = func(city int) {
		// 遍历所有邻接节点
		for neighbor := 0; neighbor < n; neighbor++ {
			// 如果与邻居城市相连且邻居未被访问
			if isConnected[city][neighbor] == 1 && !visited[neighbor] {
				visited[neighbor] = true
				// 递归访问邻居城市
				dfs(neighbor)
			}
		}
	}

	// 遍历所有城市
	for i := 0; i < n; i++ {
		// 如果当前城市未被访问，说明发现了一个新的省份
		if !visited[i] {
			visited[i] = true
			// 从当前城市开始DFS，标记所有连通的城市
			dfs(i)
			count++
		}
	}

	return count
}

// BFS

// 并查集
