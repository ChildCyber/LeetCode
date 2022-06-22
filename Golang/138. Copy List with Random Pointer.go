package leetcode

// 复制带随机指针的链表
// https://leetcode-cn.com/problems/copy-list-with-random-pointer
// 回溯 + 哈希表
type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

var cachedNode map[*RandomNode]*RandomNode

func deepCopy(node *RandomNode) *RandomNode {
	if node == nil {
		return nil
	}

	// 判断当前节点是否已复制
	if n, has := cachedNode[node]; has { // 当前节点已复制
		return n
	}

	// 当前节点还未复制
	newNode := &RandomNode{Val: node.Val} // 复制原始链表上的每个节点node创建新节点newNode
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *RandomNode) *RandomNode {
	cachedNode = map[*RandomNode]*RandomNode{}
	return deepCopy(head)
}
