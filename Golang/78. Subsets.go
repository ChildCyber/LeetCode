package leetcode

// 子集
// https://leetcode-cn.com/problems/subsets/

// 子集的定义： 原集合中任意元素组成的集合，包括空集。

// 回溯-显式（清楚写出 选与不选）
func subsets(nums []int) (ans [][]int) {
	path := []int{}

	var backtrack func(int)
	backtrack = func(start int) { // start：选到nums中索引为start的元素
		if start == len(nums) { // 已经选到nums中最后一个元素
			// 注意：这里需要创建副本，因为后续回溯会修改path
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 选当前元素
		path = append(path, nums[start])
		// 递归
		backtrack(start + 1)
		// 撤销选择（回溯）
		path = path[:len(path)-1]

		// 不选当前元素
		backtrack(start + 1)
	}

	// 从索引0开始，初始路径为空
	backtrack(0)
	return
}

// 回溯-隐式（不选当前元素的决策隐藏在for循环和递归的结构里，没进循环就是不选，递归返回时天然保留了“不选”的路径）
func subsets1(nums []int) (ans [][]int) {
	path := []int{}

	var backtrack func(int)
	backtrack = func(start int) {
		// 每次递归都记录当前路径，而不仅仅是到达末尾时
		ans = append(ans, append([]int(nil), path...))

		for i := start; i < len(nums); i++ {
			// 选这个元素
			path = append(path, nums[i])
			// 递归
			backtrack(i + 1)
			// 撤销选择（回溯）
			path = path[:len(path)-1]
		}
	}

	// 从索引0开始，初始路径为空
	backtrack(0)
	return
}
