package leetcode

// 链表最大孪生和
// https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/

// 数组
func pairSumArr(head *ListNode) int {
	// 步骤1：将链表值存储到数组中
	values := []int{}
	current := head
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	// 步骤2：使用双指针计算最大孪生和
	maxSum := 0
	left, right := 0, len(values)-1

	for left < right {
		sum := values[left] + values[right]
		if sum > maxSum {
			maxSum = sum
		}
		left++
		right--
	}

	return maxSum
}

// 快慢指针+原地反转
// O(1) 空间复杂度解法
func pairSum(head *ListNode) int {
	// 1. 使用快慢指针找到链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 反转后半部分链表
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// 3. 计算孪生和
	maxSum := 0
	p1, p2 := head, prev
	for p2 != nil {
		sum := p1.Val + p2.Val
		if sum > maxSum {
			maxSum = sum
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return maxSum
}
