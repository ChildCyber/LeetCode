package leetcode

// 最大层内元素和
// https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/

// BFS
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	maxSum := root.Val
	maxLevel := 1
	currentLevel := 1

	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		// 处理当前层的所有节点
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levelSum += node.Val

			// 将子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 更新最大层和
		if levelSum > maxSum {
			maxSum = levelSum
			maxLevel = currentLevel
		}

		currentLevel++
	}

	return maxLevel
}
