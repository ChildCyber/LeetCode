package leetcode

// 二叉树的完全性检验
// https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/
// 层序遍历
func isCompleteTree(root *TreeNode) bool {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		l := len(stack)
		for i := 0; i < l; i++ { // 该层的节点全部遍历
			node := stack[0]
			stack = stack[1:]
			// 一棵深度为h的完全二叉树树，只有在第h层才有为nil的节点
			// 节点为nil时，stack中剩余的节点也必须为nil
			if node == nil {
				for _, v := range stack {
					if v != nil {
						return false
					}
				}
				return true
			}
			stack = append(stack, node.Left, node.Right)
		}
	}
	return true
}
