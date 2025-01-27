package leetcode

// 复制带随机指针的链表
// https://leetcode-cn.com/problems/copy-list-with-random-pointer
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
	current := head
	for current != nil {
		newNode := &RandomNode{Val: current.Val}
		nodeMap[current] = newNode
		current = current.Next
	}

	// 第二次遍历：设置新节点的next和random指针
	current = head
	for current != nil {
		copyNode := nodeMap[current]

		// 设置next指针
		if current.Next != nil {
			copyNode.Next = nodeMap[current.Next]
		}
		// 设置random指针
		if current.Random != nil {
			copyNode.Random = nodeMap[current.Random]
		}

		current = current.Next
	}
	return nodeMap[head]
}
