package leetcode

// 二叉树的中序遍历
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

// 递归DFS
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	// 左 → 根 → 右
	if root != nil {
		ans = append(ans, inorderTraversal(root.Left)...)
		ans = append(ans, root.Val)
		ans = append(ans, inorderTraversal(root.Right)...)
	}
	return ans
}

// 递归-闭包
func inorderTraversalClosure(root *TreeNode) []int {
	var ans []int

	// 避免每次递归都会创建新的切片，导致多次分配内存
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			ans = append(ans, node.Val)
			inorder(node.Right)
		}
	}

	inorder(root)
	return ans
}

// 迭代-栈
func inorderTraversalIter(root *TreeNode) (ans []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil { // 到达最左叶子节点
			stack = append(stack, root) // 将当前到达的节点入栈
			root = root.Left
		}
		// 此时root为空，说明上一步的root没有左子树
		// 1. 执行左出栈。因为此时root为空，导致root.right一定为空
		// 2. 执行下一次外层while代码块，根出栈。此时root.right可能存在
		// 3a. 若root.right存在，右入栈，再出栈
		// 3b. 若root.right不存在，重复步骤2
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 出栈，记录当前节点
		ans = append(ans, root.Val)
		root = root.Right
	}
	return
}

// Morris遍历
func inorderTraversalMorris(root *TreeNode) []int {
	ans := []int{}
	cur := root

	for cur != nil {
		if cur.Left == nil {
			// 没有左子树，直接访问当前节点
			ans = append(ans, cur.Val)
			cur = cur.Right
		} else {
			// 找到左子树的最右节点
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}

			if pre.Right == nil {
				// 建立线索
				pre.Right = cur
				cur = cur.Left
			} else {
				// 恢复树结构并访问
				pre.Right = nil
				ans = append(ans, cur.Val)
				cur = cur.Right
			}
		}
	}

	return ans
}
