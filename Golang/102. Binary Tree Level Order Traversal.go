package leetcode

// 二叉树的层序遍历
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

// BFS
func levelOrder(root *TreeNode) [][]int { // 返回一个二维数组
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	// 根节点放入队列中
	q := []*TreeNode{root}
	// 遍历队列
	for len(q) > 0 { // 在while循环的每一轮中，都是将当前层的所有节点出队列，再将下一层的所有结点入队列，这样就实现了层序遍历
		// 在每一层遍历开始前，先记录队列中的节点数量levelSize（也就是这一层的结点数量），然后一口气处理完这一层的levelSize个结点
		levelSize := len(q)
		currentLevel := make([]int, 0)

		// 给BFS遍历的结果分层
		for i := 0; i < levelSize; i++ { // 变量i无实际意义，只是为了循环levelSize次
			currentLevel = append(currentLevel, q[i].Val)

			// 将下一层的节点入队
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[levelSize:] // 将当前层的所有节点出队列
		ans = append(ans, currentLevel)
	}
	return ans
}

// DFS递归
func levelOrderDFS(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	var traverse func(*TreeNode, int, [][]int)
	traverse = func(node *TreeNode, level int, res [][]int) {
		if node == nil {
			return
		}
		if len(ans) <= level {
			ans = append(ans, []int{})
		}
		// 将当前节点的值添加到对应层的列表中
		ans[level] = append(ans[level], node.Val)
		// 递归遍历左右子树，层数+1
		traverse(node.Left, level+1, ans)
		traverse(node.Right, level+1, ans)
	}

	// 递归遍历
	traverse(root, 0, ans)
	return ans
}
