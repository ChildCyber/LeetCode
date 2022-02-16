package leetcode

// 从前序与中序遍历序列构造二叉树
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 分治 + 递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	// preorder 是二叉树的前序遍历， inorder 是同一棵二叉树的中序遍历
	if len(preorder) == 0 {
		return nil
	}

	// 构造根节点
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ { // 在中序遍历中定位到根节点，可以分别知道左子树和右子树中的节点数目
		// preorder[0]：前序遍历的第一个元素为根节点
		// inorder[i]：中序遍历的第i个元素为根节点
		if inorder[i] == preorder[0] { // 获取中序遍历中的根节点的下标
			break
		}
	}
	// 前序：[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
	// 中序：[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
	// 递归地对构造出左子树和右子树，再将这两颗子树接到根节点的左右位置
	// 根节点 + 左子树的前序遍历结果长度 == 左子树的中序遍历结果长度 + 根节点
	// 前序遍历的分界线根据中序遍历的 i 来划分
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])   // 左子树的所有节点
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:]) // 右子树的所有节点
	return root
}

// 迭代
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{}
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
