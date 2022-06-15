package leetcode

// 删除链表中的节点
// https://leetcode.cn/problems/delete-node-in-a-linked-list/
// 和下一个节点交换
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
