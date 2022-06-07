package leetcode

// 二叉树的最近公共祖先
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 递归
// • 最近公共祖先的情况：
// ○ 树形一：root为p,q中的一个，这时公共祖先为root
// ○ 树形二：p,q分别在root的左右子树上（p在左子树，q在右子树；还是p在右子树，q在左子树的情况都统一放在一起考虑）这时满足p,q的最近公共祖先的结点也只有root本身
// ○ 树形三：p和q同时在root的左子树；或者p,q同时在root的右子树，这时确定最近公共祖先需要遍历子树来进行递归求解。
// • 首先要明确递归的方式，需要采用类似后序遍历的方式，先遍历左，右子树得到结果，再根据结果判断出当前树形属于前面三种讨论的哪一种，再返回相应的答案。
// • 实现细节：
// ○ 明确递归边界：root为空（此时为空树），则直接返回NULL，不存在公共祖先；root为p或q中的一个，则直接返回root;
// 这两种情况都是不需要通过递归直接能得出答案的，故直接终止递归返回即可
// ○ 用left，right分别去接收遍历左右子树得到的结点（这里递归返回值是结点指针）
// ○ 如果left和right均不为空，注意此时left,right并不是理解为子树上p,q的最近公共祖先，而是结点p或q自身的指针，这时说明p,q分别存在于root根结点的左右两端，是符合树形二的情况，直接返回root。
// ○ 如果left，right中有一个不为空，那么通过树形的分析，必然属于树形三;
// 此时left,right代表的含义就是子树上层层向上返回的最近公共祖先，最开始的调用层拿到这个结果后直接返回不空的那一个，即是答案。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归边界
	// p=root，且 q 在 root 的左或右子树中
	// q=root，且 p 在 root 的左或右子树中
	if root == nil || root == p || root == q {
		return root
	}

	// 通过递归对二叉树进行先序遍历，当遇到节点 p 或 q 时返回
	// 分解为求左子树的子问题和右子树的子问题，注意是后序遍历
	// 子问题：左右子树是否分别包含p,q
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 左右子树都有则返回root
	if left != nil && right != nil {
		return root
	}

	// 如果左右子树中有一个子问题没得到结果，则返回得到结果的子问题
	if left == nil {
		return right
	}
	return left
}

// 存储父节点
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) { // 存储祖先节点
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}

	return nil
}
