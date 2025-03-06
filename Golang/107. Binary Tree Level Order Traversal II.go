package leetcode

// 二叉树的层序遍历 II
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/

// BFS
func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			// 节点出队
			node := queue[0]
			queue = queue[1:]
			// 记录节点值
			currentLevel = append(currentLevel, node.Val)

			// 将下一层的节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, currentLevel)
	}

	// 反转结果
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}
