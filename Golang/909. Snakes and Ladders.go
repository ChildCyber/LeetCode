package leetcode

// 蛇梯棋
// https://leetcode.cn/problems/snakes-and-ladders/

// 问题本质：单源最短路径问题
// 图的抽象：
//   顶点：每个棋盘格子（1到n²）
//   边：从当前格子可以到达的下1-6个格子
//   特殊边：蛇和梯子创建了"传送边"

// BFS
// 时间复杂度：O(n²)
// 空间复杂度：O(n²)
func snakesAndLadders(board [][]int) int {
	n := len(board)
	if n == 0 {
		return 0
	}

	target := n * n

	// 坐标转换：将格子编号转换为棋盘坐标
	// 蛇形编号：奇数行从左到右，偶数行从右到左
	numToPosition := func(num int) (int, int) {
		// 转换为0-based索引
		num--
		// 计算行（从底部开始）
		row := n - 1 - num/n
		// 计算列
		col := num % n
		// 如果是偶数行（从底部数），列方向反转
		if (num/n)%2 == 1 {
			col = n - 1 - col
		}
		return row, col
	}

	// BFS初始化，起点是格子1
	queue := []int{1}

	visited := make([]bool, target+1)
	visited[1] = true

	steps := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		// 处理当前层的所有节点
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			// 检查是否到达终点
			if current == target {
				return steps
			}

			// 尝试所有可能的骰子点数（1-6）
			for dice := 1; dice <= 6; dice++ {
				nextPos := current + dice

				// 如果超出棋盘范围，跳过
				if nextPos > target {
					continue
				}

				// 转换为棋盘坐标
				row, col := numToPosition(nextPos)

				// 检查当前位置是否有蛇或梯子
				if board[row][col] != -1 {
					// 有蛇或梯子，传送到目标位置
					nextPos = board[row][col]
				}

				// 如果目标位置未被访问，加入队列
				if !visited[nextPos] {
					visited[nextPos] = true
					queue = append(queue, nextPos)
				}
			}
		}
		// 完成一层遍历，步数加1
		steps++
	}

	// 如果BFS结束仍未到达终点，返回-1
	return -1
}
