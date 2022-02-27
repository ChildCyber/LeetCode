package leetcode

// 两两交换链表中的节点
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// 两两交换链表中相邻的节点，并返回交换后链表的头节点

// 迭代
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	temp := dummyHead // 表示当前到达的节点
	for temp.Next != nil && temp.Next.Next != nil {
		// 每次交换 temp 后面的两个节点
		node1 := temp.Next
		node2 := temp.Next.Next
		// 交换节点位置，从头到尾，node1和node2位置互换，temp的next指针指向node2
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		// 重点，下次迭代时会用到，移动两个指针，到下次需要交换的节点的前一个位置
		temp = node1 // temp移动到node1
	}
	return dummyHead.Next
}

// 递归
func swapPairsRec(head *ListNode) *ListNode {
	// 递归的终止条件是链表中没有节点，或者链表中只有一个节点
	if head == nil || head.Next == nil {
		return head
	}

	// 链表中至少有两个节点
	// 原始链表的头节点变成新的链表的第二个节点，原始链表的第二个节点变成新的链表的头节点，链表中的其余节点的两两交换可以递归地实现
	one := head
	two := one.Next
	three := two.Next

	two.Next = one
	one.Next = swapPairsRec(three)
	return two
}
