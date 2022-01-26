package leetcode

// 二叉树的层序遍历
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// BFS
func levelOrder(root *TreeNode) [][]int { // 返回一个二维数组
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	// 根节点放入队列中
	q := []*TreeNode{root}
	// 遍历队列
	for len(q) > 0 { // 在 while 循环的每一轮中，都是将当前层的所有节点出队列，再将下一层的所有结点入队列，这样就实现了层序遍历
		// 在每一层遍历开始前，先记录队列中的节点数量 l（也就是这一层的结点数量），然后一口气处理完这一层的 l 个结点
		l := len(q)
		tmp := make([]int, 0)

		// 给 BFS 遍历的结果分层
		for i := 0; i < l; i++ { // 变量 i 无实际意义，只是为了循环 l 次
			tmp = append(tmp, q[i].Val)

			// 层序遍历
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[l:] // 将当前层的所有节点出队列
		res = append(res, tmp)
	}
	return res
}

// DFS
