package leetcode

import "container/heap"

// 合并K个升序链表
// https://leetcode-cn.com/problems/merge-k-sorted-lists/

// 分治 + 递归
// 借助分治的思想，把 K 个有序链表两两合并即可
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	// 分治
	mid := length / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	// 合并两条升序链表
	return mergeTwoLists23(left, right)
}

func mergeTwoLists23(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists23(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists23(l1, l2.Next)
	return l2
}

// 小根堆（优先队列）
type Item23 struct {
	node *ListNode
}
type MinHeap23 []Item23

func (h MinHeap23) Len() int            { return len(h) }
func (h MinHeap23) Less(i, j int) bool  { return h[i].node.Val < h[j].node.Val }
func (h MinHeap23) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap23) Push(x interface{}) { *h = append(*h, x.(Item23)) }
func (h *MinHeap23) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKListsHeap(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)
	for _, l := range lists {
		if l != nil {
			heap.Push(h, Item23{l})
		}
	}

	dummy := &ListNode{}
	p := dummy
	for h.Len() > 0 {
		item := heap.Pop(h).(Item23)
		p.Next = item.node
		p = p.Next
		if item.node.Next != nil {
			heap.Push(h, Item23{item.node.Next})
		}
	}
	return dummy.Next
}
