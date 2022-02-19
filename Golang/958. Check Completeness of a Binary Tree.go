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
