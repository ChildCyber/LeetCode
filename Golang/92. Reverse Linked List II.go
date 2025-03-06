package leetcode

// 反转链表 II
// https://leetcode-cn.com/problems/reverse-linked-list-ii/

// 头插法(一次遍历)
// 在需要反转的区间里，每遍历到一个节点，让这个新节点来到反转部分的起始位置
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func reverseBetween(head *ListNode, left, right int) *ListNode {
	// https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/fan-zhuan-lian-biao-ii-by-leetcode-solut-teyq/
	// 使用三个指针变量 pre、cur、next 来记录反转的过程中需要的变量：
	//   cur：指向待反转区域的第一个节点left
	//   next：指向cur的下一个节点，循环过程中，cur变化后next会变化
	//   prev：指向待反转区域的第一个节点left的前一个节点，在循环过程中不变
	dummy := &ListNode{Next: head}
	prev := dummy

	// 确定反转起点（left的前驱）
	for i := 0; i < left-1; i++ { // 从虚拟头节点prev走left-1步，来到left节点的前一个节点
		prev = prev.Next
	}

	cur := prev.Next
	// 反转left～right部分
	for i := 0; i < right-left; i++ {
		next := cur.Next
		// 重点：prev和cur保持不动，只移动next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

// 穿针引线
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func reverseBetweenHeadInsertion(head *ListNode, left, right int) *ListNode {
	// 思路：找到left和right的前驱和后继，反转left到right部分后再拼接起来
	// prev：待反转链表的首个节点的前一个节点
	// curr：待反转链表的最后节点的下一个节点
	// leftNode：待反转链表的首个节点
	// rightNode：待反转链表的最后节点
	dummy := &ListNode{Next: head}
	prev := dummy

	// 1. 从虚拟头节点 prev 走 left-1 步，来到 left 节点的前一个节点
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// 2. 从 prev 走 right-left+1 步，来到 right 节点
	rightNode := prev
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 3. 切断出一个子链表（截取链表）
	leftNode := prev.Next
	curr := rightNode.Next

	// 切断链接
	prev.Next = nil
	rightNode.Next = nil

	// 4.同lc 206，反转链表的子区间
	reverseLinkedList(leftNode)

	// 5. 接回到原来的链表中，经过反转left变为right，right变为left
	prev.Next = rightNode
	leftNode.Next = curr
	return dummy.Next
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
