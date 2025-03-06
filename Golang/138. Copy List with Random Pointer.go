package leetcode

// 复制带随机指针的链表
// https://leetcode-cn.com/problems/copy-list-with-random-pointer/

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

// 哈希表
func copyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}

	// 创建映射表：原始节点 -> 复制节点
	nodeMap := make(map[*RandomNode]*RandomNode)
	// 第一次遍历：创建所有新节点并建立映射关系
	cur := head
	for cur != nil {
		newNode := &RandomNode{Val: cur.Val}
		nodeMap[cur] = newNode
		cur = cur.Next
	}

	// 第二次遍历：设置新节点的next和random指针
	cur = head
	for cur != nil {
		copyNode := nodeMap[cur]

		// 设置next指针
		if cur.Next != nil {
			copyNode.Next = nodeMap[cur.Next]
		}
		// 设置random指针
		if cur.Random != nil {
			copyNode.Random = nodeMap[cur.Random]
		}

		cur = cur.Next
	}

	return nodeMap[head]
}
