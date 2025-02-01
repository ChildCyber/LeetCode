package leetcode

import "math"

// 最小栈
// https://leetcode-cn.com/problems/min-stack/
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor155() MinStack {
	return MinStack{
		stack:    []int{},              // 主栈：正常存储所有元素
		minStack: []int{math.MaxInt64}, // 辅助栈：栈顶元素是当前最小元素（长度比主栈多1，避免push时判断长度）
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	// 比较入栈元素是否小于当前栈顶元素
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	// 尾部元素弹出
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

// 存储差值
type MinStackDiff struct {
	stack  []int // 存储差值的栈
	minVal int   // 当前最小值
}

func ConstructorMinStackDiff() MinStackDiff {
	return MinStackDiff{
		stack:  make([]int, 0),
		minVal: 0,
	}
}

func (this *MinStackDiff) Push(val int) {
	if len(this.stack) == 0 {
		// 栈为空，直接压入0，设置minVal为val
		this.stack = append(this.stack, 0)
		this.minVal = val
	} else {
		// 计算差值并压入栈
		diff := val - this.minVal
		this.stack = append(this.stack, diff)

		// 如果新值更小，更新min
		if diff < 0 {
			this.minVal = val
		}
	}
}

func (this *MinStackDiff) Pop() {
	if len(this.stack) == 0 {
		return
	}

	// 弹出栈顶差值
	diff := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	// 如果差值为负，说明弹出的是最小值，需要更新min
	if diff < 0 {
		this.minVal = this.minVal - diff // 恢复前一个最小值
	}
}

func (this *MinStackDiff) Top() int {
	if len(this.stack) == 0 {
		return 0
	}

	diff := this.stack[len(this.stack)-1]
	if diff < 0 {
		// 如果差值为负，栈顶元素就是当前最小值
		return this.minVal
	} else {
		// 否则，栈顶元素 = min + diff
		return this.minVal + diff
	}
}

func (this *MinStackDiff) GetMin() int {
	return this.minVal
}
