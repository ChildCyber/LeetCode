package leetcode

// 斐波那契数
// https://leetcode.cn/problems/fibonacci-number/
// 递归
// 时间复杂度 O(2^n),空间复杂度 O(n)
func fib(N int) int {
	if N <= 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

// 自底向上的记忆化搜索
// 时间复杂度 O(n),空间复杂度 O(n)
func fib1(N int) int {
	if N <= 1 {
		return N
	}

	cache := map[int]int{0: 0, 1: 1}
	for i := 2; i <= N; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[N]
}

// 递推，优化版的 dp
// 时间复杂度 O(n),空间复杂度 O(1)
func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	//递推关系
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}

	return curr
}
