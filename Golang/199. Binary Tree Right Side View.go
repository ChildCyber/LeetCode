package leetcode

// 二叉树的右视图
// https://leetcode-cn.com/problems/binary-tree-right-side-view
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue) // 当前队列长度
		for i := 0; i < n; i++ {
			// 层序遍历
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[n-1].Val) // 保留最右边的节点，即队列中最后一个元素
		queue = queue[n:]                 // 该层的节点出队
	}
	return res
}

func rightSideView1(root *TreeNode) []int {
	que := make([]*TreeNode, 0)
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	que = append(que, root)
	for len(que) > 0 {
		length := len(que)

		for i := 0; i < length; i++ {
			t := que[0] // 出队
			que = que[1:]

			if t.Left != nil {
				que = append(que, t.Left)
			}
			if t.Right != nil {
				que = append(que, t.Right)
			}

			if i == length-1 {
				ans = append(ans, t.Val)
			}
		}
	}
	return ans
}
