package leetcode

// 反转链表 II
// https://leetcode-cn.com/problems/reverse-linked-list-ii/

// 一次遍历，头插法
// 在需要反转的区间里，每遍历到一个节点，让这个新节点来到反转部分的起始位置
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func reverseBetween(head *ListNode, left, right int) *ListNode {
	// https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/fan-zhuan-lian-biao-ii-by-leetcode-solut-teyq/
	// 使用三个指针变量 pre、cur、next 来记录反转的过程中需要的变量：
	//   cur：指向待反转区域的第一个节点 left
	//   next：指向 cur 的下一个节点，循环过程中，cur 变化后 next 会变化
	//   pre：指向待反转区域的第一个节点 left 的前一个节点，在循环过程中不变
	dummyNode := &ListNode{Next: head}
	pre := dummyNode
	for i := 0; i < left-1; i++ { // 从虚拟头节点 pre 走 left-1 步，来到 left 节点的前一个节点
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ { // 反转left～right部分
		// 操作步骤：
		// 先将 cur 的下一个节点记录为 next
		// 执行操作 ①：把 cur 的下一个节点指向 next 的下一个节点
		// 执行操作 ②：把 next 的下一个节点指向 pre 的下一个节点
		// 执行操作 ③：把 pre 的下一个节点指向 next
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummyNode.Next
}

// 穿针引线
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func reverseBetween1(head *ListNode, left, right int) *ListNode {
	// pre：待反转链表的首个节点的前一个节点
	// curr：待反转链表的最后节点的下一个节点
	// leftNode：待反转链表的首个节点
	// rightNode：待反转链表的最后节点
	dummyNode := &ListNode{Next: head}
	pre := dummyNode

	// 第 1 步：从虚拟头节点 pre 走 left-1 步，来到 left 节点的前一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第 2 步：从 pre 走 right-left+1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第 3 步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第 4 步：同第 206，反转链表的子区间
	reverseLinkedList(leftNode)

	// 第 5 步：接回到原来的链表中，经过反转left变为right，right变为left
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

// 反转 left 到 right 部分
func reverseLinkedList(head *ListNode) {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}
