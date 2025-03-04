package leetcode

// 删除链表的中间节点
// https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list/

// 快慢指针
func deleteMiddle(head *ListNode) *ListNode {
	// 处理空链表或只有一个节点的情况
	if head == nil || head.Next == nil {
		return nil
	}

	// 初始化指针
	slow, fast := head, head
	var prev *ListNode = nil

	// 快慢指针遍历
	for fast != nil && fast.Next != nil {
		prev = slow           // 保存慢指针的前一个节点
		slow = slow.Next      // 慢指针移动一步
		fast = fast.Next.Next // 快指针移动两步
	}

	// 删除中间节点
	prev.Next = slow.Next

	return head
}
