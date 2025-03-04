package leetcode

// 股票价格跨度
// https://leetcode.cn/problems/online-stock-span/

// 单调栈
type StockSpanner struct {
	stack [][2]int // [0]: price, [1]: span
}

func Constructor901() StockSpanner {
	return StockSpanner{
		stack: make([][2]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1 // 当前跨度至少为1（包含今天）

	// 弹出所有价格小于等于当前价格的元素，并累加它们的跨度
	for len(this.stack) > 0 && this.stack[len(this.stack)-1][0] <= price {
		span += this.stack[len(this.stack)-1][1]
		this.stack = this.stack[:len(this.stack)-1] // 弹出栈顶
	}

	// 将当前价格和跨度压入栈中
	this.stack = append(this.stack, [2]int{price, span})

	return span
}
