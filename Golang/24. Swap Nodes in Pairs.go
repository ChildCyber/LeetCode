package leetcode

// 两两交换链表中的节点
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// 迭代
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	temp := dummyHead // 头
	for temp.Next != nil && temp.Next.Next != nil {
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
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
