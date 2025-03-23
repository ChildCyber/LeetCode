package leetcode

// 合并二叉树
// https://leetcode-cn.com/problems/merge-two-binary-trees/

// 递归DFS
// 时间复杂度：O(min(m,n))
// 空间复杂度：O(min(h1,h2))
func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

// 迭代BFS
// 时间复杂度：O(min(m,n))
// 空间复杂度：O(min(m,n))
func mergeTreesBFS(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	queue := [][2]*TreeNode{{root1, root2}}

	for len(queue) > 0 {
		nodes := queue[0]
		queue = queue[1:]

		node1, node2 := nodes[0], nodes[1]

		// 合并当前节点值
		node1.Val += node2.Val

		// 处理左子树
		if node1.Left != nil && node2.Left != nil {
			queue = append(queue, [2]*TreeNode{node1.Left, node2.Left})
		} else if node1.Left == nil {
			node1.Left = node2.Left
		} // 如果node2.Left为nil，不需要做任何操作

		// 处理右子树
		if node1.Right != nil && node2.Right != nil {
			queue = append(queue, [2]*TreeNode{node1.Right, node2.Right})
		} else if node1.Right == nil {
			node1.Right = node2.Right
		} // 如果node2.Right为nil，不需要做任何操作
	}

	return root1
}
