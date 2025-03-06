package leetcode

// 两数相加
// https://leetcode-cn.com/problems/add-two-numbers/

// 模拟
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode { // 两个链表都是逆序存储数字
	dummy := &ListNode{} // 虚拟头结点
	head := dummy
	carry := 0 // 进位值

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		head.Next = &ListNode{Val: sum % 10} // 新的链表，只添加Val，不添加下个节点
		head = head.Next                     // 向后移动
		carry = sum / 10                     // carry是否大于10
	}

	return dummy.Next
}
