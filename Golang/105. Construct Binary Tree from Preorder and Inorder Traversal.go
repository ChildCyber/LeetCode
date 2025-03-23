package leetcode

// 从前序与中序遍历序列构造二叉树
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

// 分治+递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	// preorder是二叉树的前序遍历， inorder是同一棵二叉树的中序遍历
	if len(preorder) == 0 {
		return nil
	}

	// 构造根节点：前序遍历的第一个元素是根节点
	root := &TreeNode{preorder[0], nil, nil}

	// 在中序遍历中定位到根节点，可以分别知道左子树和右子树中的节点数目
	rootIndex := 0
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == preorder[0] { // 获取中序遍历中的根节点的下标
			break
		}
	}

	// 划分左右子树：前序遍历的分界线根据中序遍历的rootIndex来划分
	// 前序：[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
	// 中序：[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
	// 递归地对构造出左子树和右子树，再将这两颗子树接到根节点的左右位置
	// 根节点 + 左子树的 前序遍历 结果长度 == 左子树的 中序遍历 结果长度 + 根节点
	root.Left = buildTree(preorder[1:len(inorder[:rootIndex])+1], inorder[:rootIndex])   // 左子树的所有节点
	root.Right = buildTree(preorder[len(inorder[:rootIndex])+1:], inorder[rootIndex+1:]) // 右子树的所有节点

	return root
}

// 递归优化-传递索引
func buildTreeIndex(preorder []int, inorder []int) *TreeNode {
	// 创建中序遍历的值到索引的映射，便于快速查找根节点位置
	indexMap := make(map[int]int)
	for i, val := range inorder {
		indexMap[val] = i
	}

	// 调用递归辅助函数
	return buildTreeIndexHelper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, indexMap)
}

func buildTreeIndexHelper(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, indexMap map[int]int) *TreeNode {
	// 递归终止条件：范围无效
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	// 前序遍历的第一个元素是根节点
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的位置
	rootIndex := indexMap[rootVal]

	// 计算左子树的大小
	leftSize := rootIndex - inStart

	// 递归构建左子树
	// 左子树的前序遍历范围：从preStart+1开始，长度为leftSize
	// 左子树的中序遍历范围：从inStart开始，到rootIndex-1结束
	root.Left = buildTreeIndexHelper(preorder, preStart+1, preStart+leftSize,
		inorder, inStart, rootIndex-1, indexMap)

	// 递归构建右子树
	// 右子树的前序遍历范围：从preStart+leftSize+1开始，到preEnd结束
	// 右子树的中序遍历范围：从rootIndex+1开始，到inEnd结束
	root.Right = buildTreeIndexHelper(preorder, preStart+leftSize+1, preEnd,
		inorder, rootIndex+1, inEnd, indexMap)

	return root
}

// 迭代
func buildTreeIterative(preorder []int, inorder []int) *TreeNode {
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
