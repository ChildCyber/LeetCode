package leetcode

// 课程表
// https://leetcode-cn.com/problems/course-schedule/

// 问题本质：有向图的环检测
// 图的抽象：
//   顶点：每一门课程（0到numCourses-1）。
//   边：有向边。如果课程A是课程B的前提（即要先上A才能上B），就画一条从A到B的有向边：A→B。
// 问题建模：
//   如果能完成所有课程，就意味着可以找到一个所有课程的学习顺序。这在图论中，就等价于判断图是否存在一个拓扑排序（在这个有向图中，检测是否存在环）。

// BFS（Kahn算法）
// 时间复杂度：O(n+m)
// 空间复杂度：O(n+m)
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 思路：
	//   找到所有入度为 0 的节点（没有先修课程的课程）
	//   将这些节点移出图，并减少它们指向节点的入度
	//   重复直到没有入度为 0 的节点
	//   如果所有节点都被处理 → 无环
	//   如果还有节点剩余 → 有环

	// 构建邻接表和入度数组
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	// 初始化图和入度
	for _, pre := range prerequisites {
		course, prerequisite := pre[0], pre[1]                    // 目标课程，先修课程
		graph[prerequisite] = append(graph[prerequisite], course) // 先修课程 -> 依赖该先修课程的课程列表
		inDegree[course]++                                        // 该课程的入度+1
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
		count++ // 已经完成了一门课

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

// DFS-环检测
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	// 思路：
	//   对每个节点进行 DFS 遍历
	//   维护三种状态：未访问、访问中、已访问
	//   如果在 DFS 过程中遇到"访问中"的节点 → 检测到环

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
		for _, neighbor := range graph[node] {
			if !dfs(neighbor) {
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
