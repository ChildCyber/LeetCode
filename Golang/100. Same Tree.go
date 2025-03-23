package leetcode

// 相同的树
// https://leetcode-cn.com/problems/same-tree/

// 递归DFS
// 时间复杂度：O(min(m,n))，m和n分别是两棵树的节点数
// 空间复杂度：O(min(h1,h2))，h1和h2分别是两棵树的高度
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 基准情况：两个节点都为空
	if p == nil && q == nil {
		return true
	}

	// 基准情况：只有一个节点为空
	if p == nil || q == nil {
		return false
	}

	// 基准情况：节点值不同
	if p.Val != q.Val {
		return false
	}

	// 递归检查左右子树
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 迭代BFS
// 时间复杂度：O(min(m,n))
// 空间复杂度：O(min(m,n))
func isSameTreeIterative(p *TreeNode, q *TreeNode) bool {
	// 使用队列同时遍历两棵树
	queue := [][2]*TreeNode{{p, q}}

	for len(queue) > 0 {
		nodes := queue[0]
		queue = queue[1:]

		node1, node2 := nodes[0], nodes[1]

		// 检查两个节点是否都为空
		if node1 == nil && node2 == nil {
			continue
		}

		// 检查是否只有一个节点为空
		if node1 == nil || node2 == nil {
			return false
		}

		// 检查节点值
		if node1.Val != node2.Val {
			return false
		}

		// 将子节点加入队列
		queue = append(queue, [2]*TreeNode{node1.Left, node2.Left})
		queue = append(queue, [2]*TreeNode{node1.Right, node2.Right})
	}

	return true
}
