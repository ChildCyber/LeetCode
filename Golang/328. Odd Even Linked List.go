package leetcode

// 奇偶链表
// https://leetcode-cn.com/problems/odd-even-linked-list/

// 分离节点后合并
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head       // 奇数位置链表的头
	even := head.Next // 偶数位置链表的头
	evenHead := even  // 保存偶数链表的头，用于最后连接
	// 遍历链表，重新组织指针
	for even != nil && even.Next != nil {
		// 将奇数节点连接到奇数链表
		odd.Next = even.Next
		odd = odd.Next
		// 将偶数节点连接到偶数链表
		even.Next = odd.Next
		even = even.Next
	}
	// 将偶数链表连接在奇数链表之后
	odd.Next = evenHead

	return head
}
