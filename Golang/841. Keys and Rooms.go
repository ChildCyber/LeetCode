package leetcode

// 钥匙和房间
// https://leetcode.cn/problems/keys-and-rooms/

// BFS
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	if n == 0 {
		return true
	}

	// 记录已访问的房间
	visited := make(map[int]bool)

	// 初始化队列，从0号房间开始
	queue := []int{0}
	visited[0] = true

	// BFS遍历
	for len(queue) > 0 {
		// 取出当前房间
		currentRoom := queue[0]
		queue = queue[1:]

		// 获取当前房间的所有钥匙
		keys := rooms[currentRoom]

		// 遍历所有钥匙，访问未访问过的房间
		for _, key := range keys {
			if !visited[key] {
				visited[key] = true
				queue = append(queue, key)
			}
		}
	}

	// 检查是否访问了所有房间
	return len(visited) == n
}
