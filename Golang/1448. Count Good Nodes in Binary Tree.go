package leetcode

// 统计二叉树中的好节点
// https://leetcode.cn/problems/count-good-nodes-in-binary-tree/

// DFS
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, currentMax int) int {
		if node == nil {
			return 0
		}

		count := 0
		// 如果当前节点值大于等于路径上的最大值，则是一个好节点
		if node.Val >= currentMax {
			count = 1
			// 更新路径上的最大值
			currentMax = node.Val
		}

		// 递归处理左右子树
		leftCount := dfs(node.Left, currentMax)
		rightCount := dfs(node.Right, currentMax)

		return count + leftCount + rightCount
	}

	// 从根节点开始递归，初始最大值为根节点的值
	return dfs(root, root.Val)
}

// 栈
func goodNodesStack(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	// 使用栈存储节点和当前路径最大值
	stack := []struct {
		node     *TreeNode
		maxSoFar int
	}{{root, root.Val}}

	for len(stack) > 0 {
		// 弹出栈顶元素
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node, maxSoFar := current.node, current.maxSoFar

		// 检查是否是好节点
		if node.Val >= maxSoFar {
			count++
			maxSoFar = node.Val
		}

		// 将子节点压入栈（先右后左，保证处理顺序）
		if node.Right != nil {
			stack = append(stack, struct {
				node     *TreeNode
				maxSoFar int
			}{node.Right, maxSoFar})
		}
		if node.Left != nil {
			stack = append(stack, struct {
				node     *TreeNode
				maxSoFar int
			}{node.Left, maxSoFar})
		}
	}

	return count
}
