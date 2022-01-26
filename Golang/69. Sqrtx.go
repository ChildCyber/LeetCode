package leetcode

// x的平方根
// https://leetcode-cn.com/problems/sqrtx/
// 二分法
func mySqrt(x int) int {
	// 二分查找的下界为 0，上界可以粗略地设定为 x
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		// 在二分查找的每一步中，只需要比较中间元素 mid 的平方与 x 的大小关系，并通过比较的结果调整上下界的范围
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 牛顿迭代法
func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
