package leetcode

// 二叉树的最大深度
// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 遍历根节点的左孩子的高度和根节点右孩子的高度，取出两者的最大值再加一即为总高度
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// BFS
// 使用队列进行广度优先搜索，按层遍历节点，每遍历一层深度加1
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 使用队列进行BFS遍历
	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		// 记录当前层的节点数量
		levelSize := len(queue)

		// 处理当前层的所有节点
		for i := 0; i < levelSize; i++ {
			// 取出队列首部的节点
			node := queue[0]
			queue = queue[1:]

			// 将当前节点的子节点加入队列（如果存在）
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 处理完一层，深度加1
		depth++
	}

	return depth
}
