package leetcode

// 反转链表
// https://leetcode-cn.com/problems/reverse-linked-list/

// 双指针迭代
// 原地反转，将原链表拆分为 已反转链表 部分和 未反转链表 部分
// 遍历链表时，将当前节点的 next 指针改为指向前一个节点。由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。在更改引用之前，还需要存储后一个节点。最后返回新的头引用。
// https://leetcode-cn.com/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-shuang-zhi-zhen-di-gui-yao-mo-/
func reverseList(head *ListNode) *ListNode {
	// 定义两个指针变量 prev、curr 来记录反转的过程中需要的变量：
	//   prev：指向已反转链表的第一个节点
	//   curr：指向未反转链表的第一个节点
	// prev 最初是指向 nil，cur 指向 head，然后不断遍历 cur
	// 每次迭代到 cur，让 curr 的 next 指向 prev，实现一次局部反转
	// 局部反转完成之后，prev 和 curr 同时往前移动一个位置
	// 循环上述过程，直至 curr 到达链表尾部
	var prev *ListNode
	curr := head
	for curr != nil {
		// 局部反转
		next := curr.Next // tmp变量，记录下一个节点
		curr.Next = prev  // 向反转链表头部插入节点，第一次会插入nil
		// prev和curr指针都向后移动一个节点
		prev = curr
		curr = next
	}
	return prev
}

// 递归
// https://leetcode.cn/problems/reverse-linked-list/solution/dong-hua-yan-shi-206-fan-zhuan-lian-biao-by-user74/
func reverseListRec(head *ListNode) *ListNode {
	// 递归终止条件是当前为空，或者下一个节点为空
	if head == nil || head.Next == nil {
		return head
	}

	// newHead是最后一个节点
	newHead := reverseListRec(head.Next)
	// head 的下一个节点指向head
	head.Next.Next = head
	// 防止链表循环
	head.Next = nil
	return newHead
}

// 栈
func reverseListStack(head *ListNode) *ListNode {
	stack := make([]*ListNode, 0)
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	dummy := &ListNode{}
	newHead := dummy
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		cur.Next = nil // 防止节点间循环引用
		newHead.Next = cur
		newHead = newHead.Next
		stack = stack[:len(stack)-1]
	}
	return dummy.Next
}
