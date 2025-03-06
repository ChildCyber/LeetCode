package leetcode

// 反转链表
// https://leetcode-cn.com/problems/reverse-linked-list/

// 栈
// 空间复杂度：O(n)
func reverseListStack(head *ListNode) *ListNode {
	stack := make([]*ListNode, 0)
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	dummy := &ListNode{}
	newHead := dummy // 头节点为空节点
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		cur.Next = nil // 重点：在连接前先断开原链接，防止节点间循环引用

		newHead.Next = cur
		newHead = newHead.Next

		stack = stack[:len(stack)-1]
	}

	return dummy.Next
}

// 递归
// 空间复杂度：O(1)
func reverseListRec(head *ListNode) *ListNode {
	// 思路：先递归到链表末尾，然后在回溯的过程中逐个反转指针
	// 递归终止条件：当前为空或者下一个节点为空
	if head == nil || head.Next == nil {
		return head
	}

	// 递归到链表末尾（newHead是原先链表中的最后一个节点）
	newHead := reverseListRec(head.Next) // newHead始终指向反转后链表的头节点（原链表的尾节点）

	// 反转指针
	head.Next.Next = head // 让下一个节点指向当前节点
	head.Next = nil       // 断开当前节点原来的指向，防止链表循环

	return newHead
}

// 头插法-迭代（原地反转）
// 空间复杂度：O(1)
func reverseList(head *ListNode) *ListNode {
	// 思路：将原链表拆分为 已反转链表 部分和 未反转链表 部分
	// 定义两个指针变量 prev、cur 来记录反转的过程中需要的变量：
	//   prev：指向已反转链表的第一个节点。初始为空，因为最开始没翻转任何节点
	//   cur：指向未反转链表的第一个节点。初始为head
	var prev *ListNode
	cur := head
	for cur != nil {
		// 局部反转
		next := cur.Next // tmp变量，记录下一个节点
		cur.Next = prev  // 向反转链表头部插入节点，第一次会插入nil
		prev = cur       // prev前进
		cur = next       // cur前进
	}
	return prev
}
