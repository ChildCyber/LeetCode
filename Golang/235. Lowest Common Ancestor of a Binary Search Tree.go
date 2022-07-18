package leetcode

// 二叉搜索树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
// 递归
func lowestCommonAncestorRec(root, p, q *TreeNode) *TreeNode {
	if p == nil || q == nil || root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val { // 左子树上寻找
		return lowestCommonAncestorRec(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val { // 右子树上寻找
		return lowestCommonAncestorRec(root.Right, p, q)
	}

	return root
}

// 两次遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func getPath(root, target *TreeNode) (path []*TreeNode) { // 从二叉树上找到指定的节点，返回从根节点到指定节点的路径
	node := root
	for node != target {
		path = append(path, node)
		if target.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	path = append(path, node)
	return
}

func lowestCommonAncestorTwo(root, p, q *TreeNode) (ancestor *TreeNode) {
	pathP := getPath(root, p) // 从root上找到p
	pathQ := getPath(root, q) // 从root上找到q
	// 最近公共祖先就是从根节点到它们路径上的分岔点，也就是最后一个相同的节点
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ancestor = pathP[i] // 最后一个相同的节点
	}
	return
}

// 一次遍历
// 时间复杂度：O(n)
// 时间复杂度：O(1)
func lowestCommonAncestorOne(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	// 将两个节点放在一起遍历
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val { // p、q都小于root，向左
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val { // p、q都大于root，向右
			ancestor = ancestor.Right
		} else { // p、q一个大于一个小于，当前节点就是公共祖先
			return
		}
	}
}
