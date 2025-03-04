package leetcode

// 叶子相似的树
// https://leetcode.cn/problems/leaf-similar-trees/

// DFS
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaves1, leaves2 []int

	// 使用闭包减少函数调用和内存分配
	var getLeaves func(node *TreeNode, leaves *[]int)
	getLeaves = func(node *TreeNode, leaves *[]int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			*leaves = append(*leaves, node.Val)
			return
		}

		getLeaves(node.Left, leaves)
		getLeaves(node.Right, leaves)
	}

	getLeaves(root1, &leaves1)
	getLeaves(root2, &leaves2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}

	return true
}

// 栈
func leafSimilarDFS(root1 *TreeNode, root2 *TreeNode) bool {
	// 使用栈进行迭代DFS
	stack1, stack2 := []*TreeNode{root1}, []*TreeNode{root2}

	for len(stack1) > 0 && len(stack2) > 0 {
		// 获取下一棵树的下一片叶子
		leaf1 := getNextLeaf(&stack1)
		leaf2 := getNextLeaf(&stack2)

		// 如果叶子值不同，立即返回false
		if leaf1.Val != leaf2.Val {
			return false
		}
	}

	// 检查是否还有剩余的叶子节点
	return len(stack1) == 0 && len(stack2) == 0
}

// 辅助函数：获取下一片叶子节点
func getNextLeaf(stack *[]*TreeNode) *TreeNode {
	for len(*stack) > 0 {
		// 弹出栈顶节点
		node := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]

		// 将左右子节点压入栈（先右后左，保证左子树先处理）
		if node.Right != nil {
			*stack = append(*stack, node.Right)
		}
		if node.Left != nil {
			*stack = append(*stack, node.Left)
		}

		// 如果是叶子节点，返回它
		if node.Left == nil && node.Right == nil {
			return node
		}
	}
	return nil
}
