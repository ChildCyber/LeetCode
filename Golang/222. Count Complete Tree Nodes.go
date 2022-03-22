package leetcode

// 完全二叉树的节点个数
// https://leetcode.cn/problems/count-complete-tree-nodes/
// 层序遍历
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	curNum, nextLevelNum, ans := 1, 0, 1
	for len(queue) != 0 {
		if curNum > 0 { // 当前队列长度
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			curNum--
			queue = queue[1:] // 出队
		}

		if curNum == 0 {
			ans += nextLevelNum
			curNum = nextLevelNum
			nextLevelNum = 0
		}
	}
	return ans
}

// 递归
func countNodesRev(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countNodesRev(root.Left)
	right := countNodesRev(root.Right)
	return left + right + 1
}
