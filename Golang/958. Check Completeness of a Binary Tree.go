package leetcode

// 二叉树的完全性检验
// https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/
// 层序遍历
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ { // 该层的节点全部遍历
			node := queue[0]
			queue = queue[1:]
			// 一棵深度为h的完全二叉树树，只有在第h层才有为nil的节点
			// 当节点为nil时，queue中剩余的节点也必须为nil
			if node == nil {
				for _, v := range queue { // 层序遍历，判断queue中剩余节点是否全部为nil
					if v != nil {
						return false
					}
				}
				return true
			}
			queue = append(queue, node.Left, node.Right)
		}
	}
	return true
}

// 递归
// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/solution/er-cha-shu-de-di-gui-tao-lu-yi-rootwei-g-i9vg/
func isCompleteTreeRec(root *TreeNode) bool {
	if root == nil {
		return false
	}
	complete, _, _ := isFullTree(root)
	return complete
}

func isFullTree(root *TreeNode) (complete bool, full bool, height int) {
	// complete：当前子树是否是完全二叉树
	// full：当前子树是否满二叉树
	// height：当前子树的高度
	if root == nil {
		return true, true, 0
	}
	// 获取左右子树的信息
	leftComplete, leftFull, leftHeight := isFullTree(root.Left)
	rightComplete, rightFull, rightHeight := isFullTree(root.Right)

	// 判断是否是满二叉树
	// 左子树是满二叉树，右子树是满二叉树，左右子树一样高
	if leftFull && rightFull && leftHeight == rightHeight {
		full = true
	}

	// 判断是否是完全二叉树
	// 1.当前子树是满二叉树，一定是完全二叉树
	if full {
		complete = true
		// 2.缺口在左子树上：左树是完全二叉树，右树是满二叉树，左树高度比右树高1
	} else if leftComplete && rightFull && leftHeight == rightHeight+1 {
		complete = true
		// 3.缺口在左右子树中间：左树是满二叉树，右树是满二叉树，左子树比右子树高度高1
	} else if leftFull && rightFull && leftHeight == rightHeight+1 {
		complete = true
		// 4.缺口在右子树上：左树是满二叉树，右树是完全二叉树，左右子树高度相同
	} else if leftFull && rightComplete && leftHeight == rightHeight {
		complete = true
	} else {
		complete = false
	}
	// 当前子树的高度
	height = max(leftHeight, rightHeight) + 1

	return complete, full, height
}
