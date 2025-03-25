package leetcode

// 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum/

// 暴力
// 时间复杂度：O(nk)，n是数组长度，k是窗口大小
func maxSlidingWindowBrute(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	res := make([]int, 0, k)
	for i := 0; i <= n-k; i++ {
		maxNum := nums[i]
		for j := 1; j < k; j++ {
			if maxNum < nums[i+j] {
				maxNum = nums[i+j]
			}
		}
		res = append(res, maxNum)
	}

	return res
}

// 堆
// 大根堆维护窗口最大值，每次滑动都往堆里加新元素，同时把堆顶检查一遍，看是不是已经滑出窗口
// 元素结构：值 + 下标

// 双端队列
// 时间复杂度：O(n)
// 空间复杂度：O(k)
func maxSlidingWindow(nums []int, k int) []int {
	// 思路：队列里存下标，而不是值。保证队列里的值单调递减（队头最大）。
	n := len(nums)
	if n == 0 || n < k {
		return []int{}
	}

	deque := []int{} // 存下标，保持对应值单调递减
	res := []int{}

	// 处理每个元素：
	//   从队列尾部开始，移除所有小于当前元素的索引（因为它们不可能成为窗口最大值）
	//   将当前元素的索引加入队列尾部
	//   检查队列头部的索引是否已经滑出窗口，如果是则移除
	//   当窗口形成后（i ≥ k-1），队列头部就是当前窗口的最大值
	for i := 0; i < n; i++ {
		// 1. 移除队尾小于当前值的元素
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}

		// 2. 加入当前下标
		deque = append(deque, i)

		// 3. 移除队头已经滑出窗口的元素
		if deque[0] <= i-k {
			deque = deque[1:]
		}

		// 4. 当窗口形成时（i >= k-1），记录最大值
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}

// 分块
// 核心：预处理+快速查询
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxSlidingWindowBlock(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}

	// 划块：把整个数组nums，按照窗口大小k切成若干个块。虽然最后窗口不一定正好卡在块边界，但最多跨两个块
	left := make([]int, n)  // left[i]: 当前块从起点到 i 的最大值（从左往右扫，块头重置）
	right := make([]int, n) // right[i]: 当前块从 i 到块末尾的最大值（从右往左扫，块尾重置）

	// 从左到右预处理
	for i := 0; i < n; i++ {
		if i%k == 0 {
			left[i] = nums[i] // 块的起点
		} else {
			left[i] = max(left[i-1], nums[i])
		}
	}

	// 从右到左预处理
	for i := n - 1; i >= 0; i-- {
		if (i == n-1) || ((i+1)%k == 0) {
			right[i] = nums[i] // 块的末尾
		} else {
			right[i] = max(right[i+1], nums[i])
		}
	}

	// 查询窗口最大值
	res := make([]int, n-k+1)
	for i := 0; i <= n-k; i++ {
		res[i] = max(right[i], left[i+k-1]) // 窗口范围是[i, i+k-1]，长度 k。right[i] 涵盖了窗口的左半边，left[i+k-1] 涵盖了窗口的右半边，中间重叠部分已经被包含了
	}

	return res
}
