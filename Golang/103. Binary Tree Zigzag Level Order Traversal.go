package leetcode

// 二叉树的锯齿形层序遍历
// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
// BFS，102的变种
func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	// 根节点放入队列中
	queue := []*TreeNode{root}
	// 规定二叉树的根节点为第 0 层，如果当前层数是偶数，从左至右输出当前层的节点值，否则，从右至左输出当前层的节点值。
	for level := 0; len(queue) > 0; level++ { // 队列非空
		vals := []int{}          // 记录本层的遍历的结果
		q := queue               // 本层需要遍历的节点
		queue = nil              // 清空队列
		for _, node := range q { // 本层的节点全部出队
			vals = append(vals, node.Val) // 本层的遍历

			// 层序遍历，下一层的节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 需要把奇数层的元素翻转
		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ { // 遍历前半
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i] // 首尾互换
			}
		}
		ans = append(ans, vals)
	}
	return ans
}
