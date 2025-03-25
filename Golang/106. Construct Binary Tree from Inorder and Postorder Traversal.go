package leetcode

// 从中序与后序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

// 分治+递归
// 时间复杂度：O(n²)
func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// 后序遍历的最后一个元素是根节点
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的位置
	rootIndex := 0
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	// 中序：[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
	// 后序：[ [左子树的后序遍历结果], [右子树的后序遍历结果], 根节点 ]
	// 划分左右子树
	// 左子树的中序遍历：inorder[0:rootIndex]
	// 右子树的中序遍历：inorder[rootIndex+1:]

	// 左子树的后序遍历：postorder[0:rootIndex]
	// 右子树的后序遍历：postorder[rootIndex:len(postorder)-1]

	root.Left = buildTree(inorder[0:rootIndex], postorder[0:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])

	return root
}

// 递归优化
// 时间复杂度：O(n)
func buildTree106Map(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 || n != len(postorder) {
		return nil
	}

	// 构建中序值 -> 索引的哈希映射，O(1) 查找
	idxMap := make(map[int]int, n)
	for i, v := range inorder {
		idxMap[v] = i
	}

	postIndex := n - 1 // 指向当前后序中要作为根的元素（从后向前）

	var helper func(inL, inR int) *TreeNode
	helper = func(inL, inR int) *TreeNode {
		if inL > inR {
			return nil
		}
		if postIndex < 0 {
			return nil
		}

		rootVal := postorder[postIndex]
		postIndex--
		root := &TreeNode{Val: rootVal}

		// 在中序中找到根的位置，将区间分为左/右
		mid := idxMap[rootVal]

		// 注意：后序遍历的顺序是 left, right, root，
		// 从后向前取值，所以先构造右子树，再构造左子树
		root.Right = helper(mid+1, inR)
		root.Left = helper(inL, mid-1)
		return root
	}

	return helper(0, n-1)
}

// 迭代
func buildTree106Iterative(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	// 创建根节点
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*TreeNode{root}
	inorderIndex := len(inorder) - 1

	for i := len(postorder) - 2; i >= 0; i-- {
		postorderVal := postorder[i]
		node := stack[len(stack)-1]

		if node.Val != inorder[inorderIndex] {
			// 当前节点不是中序遍历的当前值，说明是右子节点
			node.Right = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Right)
		} else {
			// 当前节点是中序遍历的当前值，需要找到应该作为哪个节点的左子节点
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex--
			}
			node.Left = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Left)
		}
	}

	return root
}
