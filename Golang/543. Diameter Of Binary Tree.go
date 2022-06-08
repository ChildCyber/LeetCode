package leetcode

// 二叉树的直径
// https://leetcode-cn.com/problems/diameter-of-binary-tree/
// 直径定义：一棵二叉树的直径长度是任意两个结点路径长度中的最大值。路径可能穿过也可能不穿过根结点。

// 深度优先搜索
// 一条路径的长度为该路径经过的节点数减一，所以求直径（即求路径长度的最大值）等效于求路径经过节点数的最大值减一。
// 而任意一条路径均可以被看作由某个节点为起点，从其左儿子和右儿子向下遍历的路径拼接得到。
// 找出二叉树中最长的最短路径（求出长度即可）
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0 // 最大直径

	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int { // 返回该节点为根的子树的深度
		if node == nil {
			return 0
		}
		// 递归得到左右子树高度
		left := depth(node.Left)
		right := depth(node.Right)
		ans = max(left+right, ans)  // 记录目前的最大直径，最大直径=左子树高度+右子树高度
		return max(left, right) + 1 // 返回以当前节点为根的二叉树的最大深度
	}

	depth(root)
	return ans
}
