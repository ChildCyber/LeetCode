package leetcode

// 反转链表
// https://leetcode-cn.com/problems/reverse-linked-list/
// 迭代，头插法
// 原地反转，原链表拆分为已反转链表部分和未反转链表部分
func reverseList(head *ListNode) *ListNode {
	// 在遍历链表时，将当前节点的 next 指针改为指向前一个节点。由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。在更改引用之前，还需要存储后一个节点。最后返回新的头引用。
	// https://leetcode-cn.com/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-shuang-zhi-zhen-di-gui-yao-mo-/
	// 定义两个指针： prev 和 curr ；prev 在前 curr 在后（原本的链表中）。
	// 每次让 curr 的 next 指向 prev ，实现一次局部反转
	// 局部反转完成之后，prev 和 curr 同时往前移动一个位置
	// 循环上述过程，直至 curr 到达链表尾部
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next // tmp变量，记录下一个节点
		curr.Next = prev  // 局部反转，向反转链表头部插入节点
		// prev和curr指针都向后移动一个节点
		prev = curr
		curr = next
	}
	return prev
}
