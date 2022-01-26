package leetcode

import "math"

// 最小栈
// https://leetcode-cn.com/problems/min-stack/
type MinStack struct {
	stack    []int
	MinStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		MinStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	// 比较入栈元素是否小于当前栈顶元素
	top := this.MinStack[len(this.MinStack)-1]
	this.MinStack = append(this.MinStack, min(x, top))
}

func (this *MinStack) Pop() {
	// 尾部元素弹出
	this.stack = this.stack[:len(this.stack)-1]
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.MinStack)-1]
}
