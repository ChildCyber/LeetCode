package leetcode

// 二叉树的直径
// https://leetcode-cn.com/problems/diameter-of-binary-tree/
// 直径定义：一棵二叉树的直径长度是任意两个结点路径长度中的最大值
// 一条路径的长度为该路径经过的节点数减一，所以求直径（即求路径长度的最大值）等效于求路径经过节点数的最大值减一。
// 而任意一条路径均可以被看作由某个节点为起点，从其左儿子和右儿子向下遍历的路径拼接得到。
// 找出二叉树中最长的最短路径（求出长度即可）
func depth(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}

	left := depth(node.Left, ans)
	right := depth(node.Right, ans)
	*ans = max(left+right, *ans) // 最大直径=左子树高度+右子树高度（重点）
	return max(left, right) + 1  // 以当前节点为根的二叉树的最大深度
}

// 深度优先搜索
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0 // 最大直径
	depth(root, &ans)
	return ans
}
