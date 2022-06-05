package leetcode

// 合并二叉树
// https://leetcode-cn.com/problems/merge-two-binary-trees/

// DFS
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// BFS
func mergeTreesBFS(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	merged := &TreeNode{Val: t1.Val + t2.Val} // 合并后的新节点
	// 使用三个队列分别存储合并后的二叉树的节点以及两个原始二叉树的节点
	queue := []*TreeNode{merged}
	queue1 := []*TreeNode{t1}
	queue2 := []*TreeNode{t2}
	// 从根节点开始同时遍历每个二叉树，并将对应的节点进行合并
	for len(queue1) > 0 && len(queue2) > 0 {
		// 每次从每个队列中取出一个节点，判断两个原始二叉树的节点的左右子节点是否为空
		node := queue[0]
		queue = queue[1:] // pop
		node1 := queue1[0]
		queue1 = queue1[1:] // pop
		node2 := queue2[0]
		queue2 = queue2[1:] // pop
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if left1 != nil || left2 != nil {
			if left1 != nil && left2 != nil {
				left := &TreeNode{Val: left1.Val + left2.Val}
				node.Left = left
				queue = append(queue, left)
				queue1 = append(queue1, left1)
				queue2 = append(queue2, left2)
			} else if left1 != nil {
				node.Left = left1
			} else { // left2 != nil
				node.Left = left2
			}
		}
		if right1 != nil || right2 != nil {
			if right1 != nil && right2 != nil {
				right := &TreeNode{Val: right1.Val + right2.Val}
				node.Right = right
				queue = append(queue, right)
				queue1 = append(queue1, right1)
				queue2 = append(queue2, right2)
			} else if right1 != nil {
				node.Right = right1
			} else {
				node.Right = right2
			}
		}
	}
	return merged
}
