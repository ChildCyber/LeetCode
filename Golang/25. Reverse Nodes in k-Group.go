package leetcode

// K 个一组翻转链表
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

// 递归
func reverseKGroupRecursive(head *ListNode, k int) *ListNode {
	// 1. 检查剩余长度是否 >= k
	// endTail将会指向当前组的最后一个节点（原 tail）
	endTail := head
	for i := 0; i < k; i++ {
		if endTail == nil {
			return head
		}
		endTail = endTail.Next
	}
	// 循环结束后，endTail指向的是下一组的头部

	// 2. 递归下一组
	// nextHead是下一组的头部，即当前组反转后的新尾巴应该指向的位置
	nextHead := reverseKGroupRecursive(endTail, k)

	// 3. 原地反转当前k个节点 [head, endTail)
	// prev: 新链表头部 (即当前组反转后的新尾巴)
	prev := nextHead // 注意：prev直接指向下一组反转后的新头部，完成连接
	curr := head     // curr: 当前处理的节点

	// 只需要反转k次
	for i := 0; i < k; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// 此时prev已经指向了当前组反转后的新头部（原tail）
	return prev
}

// 模拟迭代法（双指针+原地翻转）
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

		next := tail.Next                              // 记录反转后需要重新连接的下个节点
		newHead, newTail := reverseSegment(head, tail) // 反转head到tail部分的链表
		// 把子链表重新接回原链表
		pre.Next = newHead
		newTail.Next = next // 可省略

		// pre和head向后移动k次
		pre = newTail // 指向本次反转链表部分尾部
		head = next   // 下次反转链表部分头部
	}

	return dummy.Next
}

// 原地反转链表，将原链表拆分为 已反转链表 部分和 未反转链表 部分
func reverseSegment(head, tail *ListNode) (*ListNode, *ListNode) {
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

// 栈
func reverseKGroupStack(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil {
		stack := []*ListNode{}
		tmp := head
		// 收集 k 个节点
		for i := 0; i < k && tmp != nil; i++ {
			stack = append(stack, tmp)
			tmp = tmp.Next
		}
		if len(stack) < k {
			pre.Next = head
			break
		}
		// 出栈翻转
		for i := k - 1; i >= 0; i-- {
			pre.Next = stack[i]
			pre = pre.Next
		}
		// 接回链表
		pre.Next = tmp
		head = tmp
	}
	return dummy.Next
}
