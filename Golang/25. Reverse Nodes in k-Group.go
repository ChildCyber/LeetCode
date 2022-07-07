package leetcode

// K 个一组翻转链表
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

// 递归
func reverseGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse25(head, node)
	head.Next = reverseGroup(node, k)
	return newHead
}

func reverse25(head *ListNode, tail *ListNode) *ListNode {
	prev := tail
	for head != tail {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// 模拟
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 链表分为 已翻转部分+待翻转部分+未翻转部分
	// head：反转部分的链表的头部，永远在pre的后一个节点上
	// tail：反转部分的链表的尾部
	// pre：记录反转部分的链表的头部的前一个节点
	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil {
		tail := pre // tail每次循环会重新赋值

		// 判断未翻转部分长度是否大于等于k
		for i := 0; i < k; i++ {
			tail = tail.Next // 经过k次移动，得到需要反转的链表尾部tail
			if tail == nil { // 长度小于k，不反转直接返回
				return dummy.Next
			}
		}

		nex := tail.Next                   // 记录反转后需要重新连接的下个节点
		head, tail = myReverse(head, tail) // 反转head到tail部分的链表
		// 把子链表重新接回原链表
		pre.Next = head
		tail.Next = nex // 可省略

		// pre和head向后移动k次
		pre = tail       // 指向本次反转链表部分尾部
		head = tail.Next // 下次反转链表部分头部
	}

	return dummy.Next
}

// 原地反转链表，将原链表拆分为 已反转链表 部分和 未反转链表 部分
func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next // 避免再连接tail的下一个节点
	curr := head
	for prev != tail { // tail的指向一直不变，当prev==tail时跳出循环
		// 局部反转
		next := curr.Next
		curr.Next = prev
		// prev和curr指针向后移动
		prev = curr
		curr = next
	}
	return tail, head // 反转链表后，首尾节点互换
}
