package leetcode

// 课程表 II
// https://leetcode.cn/problems/course-schedule-ii/

// 问题本质：检测有向图是否存在环；如果无环，输出拓扑排序
// 图的抽象：
//   顶点：每一门课程（0到numCourses-1）。
//   边：有向边。如果课程A是课程B的前提（即要先上A才能上B），就画一条从A到B的有向边：A→B。
// 问题建模：
//   如果能完成所有课程，就意味着可以找到一个所有课程的学习顺序。这在图论中，就等价于判断图是否存在一个拓扑排序（在这个有向图中，检测是否存在环）。

// BFS（Kahn算法）
// 时间复杂度：O(n+m)
// 空间复杂度：O(n+m)
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 1. 初始化入度数组和邻接表
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	// 2. 构建图
	for _, pre := range prerequisites {
		course, prerequisite := pre[0], pre[1]                    // 学习course之前需要先学习prerequisite
		graph[prerequisite] = append(graph[prerequisite], course) // prerequisite -> course
		inDegree[course]++
	}

	// 3. 找到所有入度为0的节点（没有先修要求的课程）
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 4. BFS处理 + 记录拓扑顺序
	ans := make([]int, 0, numCourses)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		ans = append(ans, course) // 任意可行解

		// 减少后继课程的入度
		for _, neighbor := range graph[course] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 5. 判断结果
	if len(ans) == numCourses { // 结果数组长度 = 总课程数 → 无环
		return ans
	}
	return []int{} // 有环
}
