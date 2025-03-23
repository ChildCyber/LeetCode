package leetcode

// 二叉搜索树中的插入操作
// https://leetcode.cn/problems/insert-into-a-binary-search-tree/

// 递归DFS
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 基准情况：找到插入位置
	if root == nil {
		return &TreeNode{Val: val}
	}

	// 根据BST性质决定插入方向
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

// 迭代
func insertIntoBSTIterative(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}

	if root == nil {
		return newNode
	}

	var parent *TreeNode // anchor指针：始终指向最后一个非空的节点（即未来新节点的父节点）
	cur := root          // 进行遍历和查找空位

	// 找到插入位置的父节点
	for cur != nil { // BST插入要求在叶子节点的下方找到一个空位来挂载新节点
		parent = cur
		if val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	// 在父节点的适当位置插入新节点
	if val < parent.Val {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	return root
}
