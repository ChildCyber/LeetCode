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

// 哈希表
func hasCycleHash(head *ListNode) bool {
	// 遍历链表中的每个节点，并将它记录下来；一旦遇到了此前遍历过的节点，就可以判定链表中存在环
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}
