package leetcode

// 环形链表 II
// https://leetcode-cn.com/problems/linked-list-cycle-ii/
// 快慢指针
// 第 141 题的加强版。在判断是否有环的基础上，还需要输出环的第一个点
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next      // 慢
		if fast.Next == nil { // 链表尽头
			return nil
		}
		fast = fast.Next.Next // 快
		if fast == slow {     // 快慢指针相遇
			p := head       // 快指针回到链表的头部
			for p != slow { // 两指针同时向前走
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// 哈希表
func detectCycleHash(head *ListNode) *ListNode {
	// 遍历链表中的每个节点，并将它记录下来；一旦遇到了此前遍历过的节点，就可以判定链表中存在环
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}
