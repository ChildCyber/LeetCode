package leetcode

// 排序链表
// https://leetcode-cn.com/problems/sort-list/
// 自顶向下归并排序
func merge148(head1, head2 *ListNode) *ListNode { // 合并两个排序后的子链表
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sort148(head, tail *ListNode) *ListNode { // 递归
	// 1. 找到链表的中点，以中点为分界，将链表拆分成两个子链表。寻找链表的中点可以使用快慢指针的做法，快指针每次移动 2 步，慢指针每次移动 1 步，当快指针到达链表末尾时，慢指针指向的链表节点即为链表的中点。
	// 2. 对两个子链表分别排序。
	// 3. 将两个排序后的子链表合并，得到完整的排序后的链表。

	// 递归的终止条件是链表的节点个数小于或等于 1
	// 当链表为空或者链表只包含 1 个节点时，不需要对链表进行拆分和排序。
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge148(sort148(head, mid), sort148(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort148(head, nil)
}
