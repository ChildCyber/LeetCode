package leetcode

// 求根节点到叶节点数字之和
// https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
// DFS
func sumNumbers(root *TreeNode) int {
	res := 0 // 和
	dfs129(root, 0, &res)
	return res
}

func dfs129(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	// 当前节点和子树的和
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += sum
		return
	}
	dfs129(root.Left, sum, res)
	dfs129(root.Right, sum, res)
}
