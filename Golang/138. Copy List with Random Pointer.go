package leetcode

// 复制带随机指针的链表
// https://leetcode-cn.com/problems/copy-list-with-random-pointer
// 回溯 + 哈希表
var cachedNode map[*RandomNode]*RandomNode

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func deepCopy(node *RandomNode) *RandomNode {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &RandomNode{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *RandomNode) *RandomNode {
	cachedNode = map[*RandomNode]*RandomNode{}
	return deepCopy(head)
}
