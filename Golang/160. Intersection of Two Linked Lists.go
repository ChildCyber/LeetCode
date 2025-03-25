package leetcode

// 相交链表
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

// 双指针
// 给定的 2 个链表的⻓度如果一样⻓，都从头往后扫即可。
// 如果不一样⻓，需要先“拼成”一样⻓。把 B 拼接到 A 后面，把 A 拼接到 B 后面。这样 2 个链表的⻓度都是 A + B。
// 再依次扫描比较 2 个链表的结点是否相同。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 都不为空，两个链表才可能相交
	if headA == nil || headB == nil {
		return nil
	}
	// 当链表 headA 和 headB 都不为空时，创建两个指针 a 和 b，初始时分别指向两个链表的头节点 headA和 headB，然后将两个指针依次遍历两个链表的每个节点。
	a, b := headA, headB
	for a != b { // 退出循环
		if a == nil { // A指针为空，指针移到链表headB的头节点
			a = headB
		} else {
			a = a.Next
		}

		if b == nil { // B指针为空，指针移到链表headA的头节点
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

// 哈希表
// 集合存储链表节点：先遍历链表headA，并将链表headA中的每个节点加入集合中。
// 遍历链表headB，判断该节点是否在集合中：不在，继续遍历下个节点；在，返回该节点。
// headB中所有节点都不在集合中，则两链表不相交。
func getIntersectionNodeSet(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}
