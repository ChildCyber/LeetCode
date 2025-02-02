package leetcode

// 扁平化多级双向链表
// https://leetcode.cn/problems/flatten-a-multilevel-doubly-linked-list/
/**
 * Definition for a Node.
 * type CNode struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

type CNode struct {
	Val   int
	Prev  *CNode
	Next  *CNode
	Child *CNode
}

// 递归
func flatten430(root *CNode) *CNode {
	if root == nil {
		return nil
	}

	// 递归处理整个链表
	flatten430Helper(root)
	return root
}

// flatten430Helper 递归展平多级链表，返回展平后的尾节点
func flatten430Helper(head *CNode) *CNode {
	current := head
	var tail *CNode // 记录当前链表的尾节点

	for current != nil {
		next := current.Next // 保存下一个节点

		// 如果当前节点有子链表
		if current.Child != nil {
			// 递归展平子链表，获取子链表的头尾节点
			childHead := current.Child
			childTail := flatten430Helper(childHead)

			// 将当前节点与子链表连接
			current.Next = childHead
			childHead.Prev = current

			// 将子链表与下一个节点连接
			if next != nil {
				childTail.Next = next
				next.Prev = childTail
			}

			// 清空child指针
			current.Child = nil

			// 更新尾节点为子链表的尾节点
			tail = childTail
		} else {
			// 没有子链表，尾节点就是当前节点
			tail = current
		}

		// 移动到下一个节点
		current = next
	}

	return tail
}

// 迭代+栈
func flatten430Stack(root *CNode) *CNode {
	if root == nil {
		return nil
	}

	stack := []*CNode{root}
	var prev *CNode = nil

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 处理前驱指针
		if prev != nil {
			prev.Next = curr
			curr.Prev = prev
		}

		// 遍历的顺序是先当前节点，然后处理子节点，然后处理下一个节点
		// 所以应该先压入下一个节点，再压入子节点。这样栈顶是子节点，先处理子节点
		// 将next和child按顺序入栈（注意顺序）
		if curr.Next != nil {
			stack = append(stack, curr.Next)
		}
		if curr.Child != nil {
			stack = append(stack, curr.Child)
			curr.Child = nil
		}

		prev = curr
	}

	return root
}
