package leetcode

// 重排链表
// https://leetcode-cn.com/problems/reorder-list/
// 数组
// 数组存储该链表，然后利用数组可以下标访问的特点，直接按顺序访问指定元素，重建该链表即可。
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// 数组存储链表
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	i, j := 0, len(nodes)-1 // 首，尾
	// 反转
	for i < j { // 重点
		nodes[i].Next = nodes[j]
		i++
		if i == j { // 首尾相交
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil // 重点
}

func reorderList1(head *ListNode) {
	if head == nil {
		return
	}

	// 数组存储链表
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	length := len(nodes)
	cur := head
	last := head
	for i := 0; i < length/2; i++ {
		tmp := &ListNode{Val: nodes[length-1-i].Val, Next: cur.Next} // i对应的尾部节点，Next指向当前节点的下个节点
		cur.Next = tmp
		cur = tmp.Next // cur向后移动
		last = tmp     // 链表尾部
	}
	if length%2 == 0 {
		last.Next = nil
	} else {
		cur.Next = nil
	}
}

// 寻找链表中点 + 链表逆序 + 合并链表
