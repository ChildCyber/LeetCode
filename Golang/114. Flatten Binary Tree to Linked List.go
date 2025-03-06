package leetcode

// 二叉树展开为链表
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

// 先序遍历-非原地
func flatten(root *TreeNode) {
	list := preorderTraversal114(root)
	for i := 1; i < len(list); i++ { // 铺开二叉树
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr // 更新每个节点的左右子节点的信息，将二叉树展开为单链表
	}
}

func preorderTraversal114(root *TreeNode) []*TreeNode {
	list := make([]*TreeNode, 0)
	// 先序遍历
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal114(root.Left)...)
		list = append(list, preorderTraversal114(root.Right)...)
	}
	return list
}

// 递归-原地
func flattenRec(root *TreeNode) {
	if root == nil {
		return
	}

	// 将左子树展开为链表
	flattenRec(root.Left)
	// 将右子树展开为链表
	flattenRec(root.Right)
	// 将左子树迁移到右子树
	right := root.Right
	root.Left, root.Right = nil, root.Left
	// 找到新的右子树末端
	lastNode := root
	for lastNode.Right != nil {
		lastNode = lastNode.Right
	}
	// 将原右子树拼接到新右子树末端
	lastNode.Right = right
}
