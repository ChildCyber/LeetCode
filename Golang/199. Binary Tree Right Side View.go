package leetcode

// 二叉树的右视图
// https://leetcode-cn.com/problems/binary-tree-right-side-view/

// 层序遍历BFS
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	// 使用队列进行层序遍历
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// 对于每一层，记录该层的节点数量
		levelSize := len(queue)
		// 遍历当前层的所有节点，但只记录最后一个节点的值
		for i := 0; i < levelSize; i++ {
			// 将下一层的节点入队
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		ans = append(ans, queue[levelSize-1].Val) // 保留最右边的节点，即队列中最后一个元素
		queue = queue[levelSize:]                 // 该层的节点出队
	}
	return ans
}
