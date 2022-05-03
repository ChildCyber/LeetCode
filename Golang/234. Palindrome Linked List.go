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
func isPalindromeLL(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	res := true
	// 寻找中间节点
	p1 := head // 快慢指针，p1为中间节点
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	// 反转链表后半部分
	preMiddle := p1
	preCurrent := p1.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}
	// 扫描表，判断是否回文
	p1 = head
	p2 = preMiddle.Next
	for p1 != preMiddle {
		if p1.Val == p2.Val {
			p1 = p1.Next
			p2 = p2.Next
		} else {
			res = false
			break
		}
	}
	if p1 == preMiddle {
		if p2 != nil && p1.Val != p2.Val {
			return false
		}
	}
	return res
}
