package leetcode

// 奇偶链表
// https://leetcode-cn.com/problems/odd-even-linked-list/
// 分离节点后合并
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 维护两个指针 head 和 evenHead 分别指向奇数链表头节点和偶数链表头节点
	evenHead := head.Next
	// 维护两个指针 odd 和 even 分别指向奇数链表尾节点和偶数链表尾节点
	odd := head
	even := evenHead
	for even != nil && even.Next != nil { // 重点
		// 更新奇数节点
		odd.Next = even.Next // 构建奇数链表
		odd = odd.Next       // 奇数节点向后移动
		// 更新偶数节点
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead // 将偶数链表连接在奇数链表之后

	return head
}
