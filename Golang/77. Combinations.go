package leetcode

// 组合
// https://leetcode-cn.com/problems/combinations/
// 回溯
func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	// 从 1 开始是题目的设定
	generateCombinations(n, k, 1, c, &res)
	return res
}

func generateCombinations(n, k, start int, path []int, res *[][]int) {
	// start：搜索起点，表示在区间 [start, n] 里选出若干个数的组合
	// 递归终止条件是：path 的长度等于 k
	if len(path) == k {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	// for i := start; i <= n; i++ {
	// 剪枝，搜索起点上界， n - (k - c.size()) + 1
	// 搜索起点的上界 + 接下来要选择的元素个数 - 1 = n => 搜索起点的上界 + 接下来要选择的元素个数 - 1 = n
	for i := start; i <= n-(k-len(path))+1; i++ { // 从下标start开始
		// 向路径变量里添加一个数
		path = append(path, i)
		// 下一轮搜索，设置的搜索起点要加 1，因为组合数中不允许出现重复的元素
		generateCombinations(n, k, i+1, path, res)
		// 回溯
		path = path[:len(path)-1]
	}
	return
}
