package leetcode

import (
	"strconv"
	"strings"
)

// 二叉树的序列化与反序列化
// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

// DFS
type Codec struct{}

func Constructor297() (_ Codec) {
	return
}

// 前序遍历DFS
func (Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return sb.String()
}

func (Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")

	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}

	return build()
}
