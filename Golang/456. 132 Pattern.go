package leetcode

import "math/rand"

// 132模式
// https://leetcode-cn.com/problems/132-pattern
// 枚举3
type node struct {
	ch       [2]*node
	priority int
	val      int
	cnt      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

type treap struct {
	root *node
}

func (t *treap) _put(o *node, val int) *node {
	if o == nil {
		return &node{priority: rand.Int(), val: val, cnt: 1}
	}
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._put(o.ch[d], val)
		if o.ch[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.cnt++
	}
	return o
}

func (t *treap) put(val int) {
	t.root = t._put(t.root, val)
}

func (t *treap) _delete(o *node, val int) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], val)
		return o
	}
	if o.cnt > 1 {
		o.cnt--
		return o
	}
	if o.ch[1] == nil {
		return o.ch[0]
	}
	if o.ch[0] == nil {
		return o.ch[1]
	}
	d := 0
	if o.ch[0].priority > o.ch[1].priority {
		d = 1
	}
	o = o.rotate(d)
	o.ch[d] = t._delete(o.ch[d], val)
	return o
}

func (t *treap) delete(val int) {
	t.root = t._delete(t.root, val)
}

func (t *treap) upperBound(val int) (ub *node) {
	for o := t.root; o != nil; {
		if o.cmp(val) == 0 {
			ub = o
			o = o.ch[0]
		} else {
			o = o.ch[1]
		}
	}
	return
}

func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	leftMin := nums[0]
	rights := &treap{}
	for _, v := range nums[2:] {
		rights.put(v)
	}

	for j := 1; j < n-1; j++ {
		if nums[j] > leftMin {
			ub := rights.upperBound(leftMin)
			if ub != nil && ub.val < nums[j] {
				return true
			}
		} else {
			leftMin = nums[j]
		}
		rights.delete(nums[j+1])
	}

	return false
}
