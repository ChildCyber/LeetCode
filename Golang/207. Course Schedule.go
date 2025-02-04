package leetcode

// 课程表
// https://leetcode-cn.com/problems/course-schedule/
// 拓扑排序问题

// DFS-环检测
// 将深度优先搜索的流程与拓扑排序的求解联系起来，用一个栈来存储所有已经搜索完成的节点
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	// 思路：沿着依赖链深入探索，如果探索过程中回到正在探索的任务，就说明有环
	// 1. 建邻接表
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		a, b := pre[0], pre[1] // 先修b才能学a, b->a
		graph[b] = append(graph[b], a)
	}

	// 0 = 未访问, 1 = 当前递归栈中, 2 = 已访问完毕
	visited := make([]int, numCourses)

	var dfs func(int) bool
	dfs = func(node int) bool {
		if visited[node] == 1 {
			// 在当前递归栈里再次遇到，说明有环
			return false
		}
		if visited[node] == 2 {
			// 已经访问完毕，直接返回
			return true
		}

		visited[node] = 1 // 标记当前路径
		// 递归访问所有邻居
		for _, nei := range graph[node] {
			if !dfs(nei) {
				return false
			}
		}
		visited[node] = 2 // 标记当前节点为已访问完成
		return true
	}

	// 2. 遍历所有节点
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if !dfs(i) {
				return false
			}
		}
	}
	return true
}

// BFS-拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构建邻接表和入度数组
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	// 初始化图和入度
	for _, v := range prerequisites {
		to, from := v[0], v[1]                // 目标课程，先修课程
		graph[from] = append(graph[from], to) // 先修课程 -> 依赖该先修课程的课程列表
		inDegree[to]++                        // 该课程的入度+1
	}

	// 找到所有入度为0的节点（可以直接学习的课程）
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS拓扑排序
	count := 0 // 记录已学习的课程数量
	for len(queue) > 0 {
		// 取出当前可以学习的课程
		course := queue[0]
		queue = queue[1:]
		count++

		// 学完当前课程后，更新依赖它的课程的入度
		for _, neighbor := range graph[course] {
			inDegree[neighbor]--
			// 如果某个课程的入度变为0，说明可以学习了
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 如果所有课程都能学习，返回true，否则返回false
	return count == numCourses
}
