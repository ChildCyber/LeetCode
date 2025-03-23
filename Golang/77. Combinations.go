package leetcode

// 组合
// https://leetcode-cn.com/problems/combinations/

// 回溯
func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}

	var ans [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// 递归终止条件是path的长度等于k
		if len(path) == k {
			// 创建路径的副本并添加到结果中
			tmp := make([]int, k)
			copy(tmp, path)
			ans = append(ans, tmp)
			return // 只是结束当前递归，不是提前终止搜索
		}

		// 剪枝：当剩余可选元素不足以填满所需数量时，立即停止
		//   当前已选len(path)个元素的数量，还需要选k-len(path)个元素
		// i:=start：不可回头，每次都从上一次选择的后面开始
		for i := start; i <= n-(k-len(path))+1; i++ {
			// 做选择
			path = append(path, i)
			// 递归：i+1，只能向前
			backtrack(i + 1)
			// 撤销选择（回溯）
			path = path[:len(path)-1]
		}
	}

	// 从1开始是题目的设定
	backtrack(1)

	return ans
}
