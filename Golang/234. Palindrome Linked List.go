package leetcode

// 回文链表
// https://leetcode-cn.com/problems/palindrome-linked-list/
// 将值复制到数组中后用双指针法
func isPalindromeL(head *ListNode) bool {
	// 链表复制到数组
	vals := []int{}
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	// 双指针
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}

// 递归
func isPalindromeRec(head *ListNode) bool {
	frontPointer := head

	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}

// 快慢指针
// 找到中间结点，反转中间节点后面到结尾的所有节点。判断头节点开始的节点和中间节点往后开始的节点是否相等
func isPalindromeFastSlowP(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 1. 使用快慢指针找到链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 反转后半部分链表
	var prev *ListNode
	current := slow
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// 3. 比较前半部分和反转后的后半部分
	firstHalf, secondHalf := head, prev
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return true
}
