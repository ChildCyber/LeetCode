package leetcode

// 环形链表
// https://leetcode-cn.com/problems/linked-list-cycle/
// 快慢指针
func hasCycle(head *ListNode) bool {
	// 同一起点出发
	fast := head
	slow := head
	for slow != nil && fast != nil && fast.Next != nil { // 头节点不为空或只有一个节点
		fast = fast.Next.Next // 移动指针
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
