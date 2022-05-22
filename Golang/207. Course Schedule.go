package leetcode

// 课程表
// https://leetcode-cn.com/problems/course-schedule/
// 拓扑排序问题

// DFS，将深度优先搜索的流程与拓扑排序的求解联系起来，用一个栈来存储所有已经搜索完成的节点。
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// BFS
func canFinish(n int, pre [][]int) bool {
	in := make([]int, n)
	frees := make([][]int, n)
	next := make([]int, 0, n)

	for _, v := range pre {
		in[v[0]]++
		frees[v[1]] = append(frees[v[1]], v[0])
	}

	for i := 0; i < n; i++ {
		if in[i] == 0 {
			next = append(next, i)
		}
	}

	for i := 0; i != len(next); i++ {
		c := next[i]
		v := frees[c]
		for _, vv := range v {
			in[vv]--
			if in[vv] == 0 {
				next = append(next, vv)
			}
		}
	}
	return len(next) == n
}
