package leetcode

// 两数相加
// https://leetcode-cn.com/problems/add-two-numbers/
// 模拟
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode { // 两个链表都是逆序存储数字
	head := &ListNode{Val: 0} // 虚拟头结点，虚拟头结点的 Next 指向真正的 head
	// 当前两个链表处相应位置的数字为 node1, node2，进位值为 carry，curr为当前节点
	node1, node2, carry, curr := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			node1 = 0
		} else {
			node1 = l1.Val
			l1 = l1.Next // l1非空，向后移动
		}

		if l2 == nil {
			node2 = 0
		} else {
			node2 = l2.Val
			l2 = l2.Next // 非空，向后移动
		}

		curr.Next = &ListNode{Val: (node1 + node2 + carry) % 10} // 新的链表，只添加Val，不添加下个节点
		curr = curr.Next                                         // 向后移动
		carry = (node1 + node2 + carry) / 10                     // carry是否大于10
	}
	return head.Next
}
