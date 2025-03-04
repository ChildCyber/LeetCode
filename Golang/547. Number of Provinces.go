package leetcode

// 省份数量
// https://leetcode.cn/problems/number-of-provinces/

// DFS
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n) // 记录城市是否被访问过
	count := 0

	var dfs func(int)
	dfs = func(i int) { // DFS函数：从当前城市出发，标记所有连通的城市
		for j := 0; j < n; j++ {
			// 如果与邻居城市相连且邻居未被访问
			if isConnected[i][j] == 1 && !visited[j] {
				visited[j] = true
				// 递归访问邻居城市
				dfs(j)
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
