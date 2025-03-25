package leetcode

// 两两交换链表中的节点
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// 两两交换链表中相邻的节点，并返回交换后链表的头节点
// 本质上是K=2的K组反转问题（lc 25）

// 递归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func swapPairsRecursive(head *ListNode) *ListNode {
	// 递归的终止条件是链表中没有节点，或者链表中只有一个节点
	if head == nil || head.Next == nil {
		return head
	}

	// 链表中至少有两个节点：原始链表的头节点变成新的链表的第二个节点，原始链表的第二个节点变成新的链表的头节点，链表中的其余节点的两两交换可以递归地实现
	// 定义要交换的两个节点
	first := head
	second := head.Next

	// 递归交换后续节点
	first.Next = swapPairsRecursive(second.Next)

	// 交换当前两个节点
	second.Next = first

	return second
}

// 迭代
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head} // 哑节点：避免处理头节点交换的特殊情况
	prev := dummy                  // 前驱指针：指向当前要交换的两个节点的前一个节点

	// 还有一对节点可以交换
	for prev.Next != nil && prev.Next.Next != nil {
		// 获取要交换的两个节点
		first := prev.Next       // 当前要交换的第一个节点
		second := prev.Next.Next // 当前要交换的第二个节点

		// 交换节点位置
		first.Next = second.Next
		second.Next = first
		prev.Next = second

		// 移动prev到下一对的前一个节点
		prev = first
	}

	return dummy.Next
}
