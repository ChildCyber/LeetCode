package leetcode

// 子集
// https://leetcode-cn.com/problems/subsets/

// 回溯-显式（清楚写出选与不选）
func subsets(nums []int) (ans [][]int) {
	path := []int{}

	var dfs func(int)
	dfs = func(cur int) { // cur：选到nums中下标为cur的元素
		if cur == len(nums) { // 已经选到nums中最后一个元素
			// 注意：这里需要创建副本，因为后续回溯会修改path
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 选当前元素
		path = append(path, nums[cur])
		dfs(cur + 1)
		path = path[:len(path)-1]

		// 不选当前元素
		dfs(cur + 1)
	}

	// 从索引0开始，初始路径为空
	dfs(0)
	return
}

// 回溯-隐式（用 for 去选，没进循环就是不选，递归返回时天然保留了“不选”的路径）
func subsets1(nums []int) (ans [][]int) {
	path := []int{}

	var dfs func(int)
	dfs = func(cur int) {
		// 每次递归都记录当前路径，而不仅仅是到达末尾时
		ans = append(ans, append([]int(nil), path...))

		for i := cur; i < len(nums); i++ {
			// 选这个元素
			path = append(path, nums[i])
			dfs(i + 1)
			// 撤销选择
			path = path[:len(path)-1]
		}
	}

	// 从索引0开始，初始路径为空
	dfs(0)
	return
}
