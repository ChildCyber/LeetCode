package leetcode

// 反转链表 II
// https://leetcode-cn.com/problems/reverse-linked-list-ii/
// 头插法
// 整体思想是：在需要反转的区间里，每遍历到一个节点，让这个新节点来到反转部分的起始位置
func reverseBetween(head *ListNode, left, right int) *ListNode {
	// 使用三个指针变量 pre、curr、next 来记录反转的过程中需要的变量：
	//    cur：指向待反转区域的第一个节点 left；
	//    next：永远指向 curr 的下一个节点，循环过程中，curr 变化以后 next 会变化；
	//    pre：永远指向待反转区域的第一个节点 left 的前一个节点，在循环过程中不变。
	dummyNode := &ListNode{Next: head}
	pre := dummyNode
	for i := 0; i < left-1; i++ { // 循环条件
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ { // 循环条件
		// 操作步骤：
		// 先将 curr 的下一个节点记录为 next；
		// 执行操作 ①：把 curr 的下一个节点指向 next 的下一个节点；
		// 执行操作 ②：把 next 的下一个节点指向 pre 的下一个节点；
		// 执行操作 ③：把 pre 的下一个节点指向 next。
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
