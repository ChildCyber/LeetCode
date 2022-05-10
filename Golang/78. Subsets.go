package leetcode

// 子集
// https://leetcode-cn.com/problems/subsets/
// 回溯
func subsets(nums []int) [][]int {
	c, ans := []int{}, [][]int{}
	for k := 0; k <= len(nums); k++ { // 子集的长度为[0:len(nums)+1]
		generateSubsets(nums, k, 0, c, &ans) // 假设子集的长度为k
	}
	return ans
}

func generateSubsets(nums []int, k, start int, c []int, ans *[][]int) {
	// k：子集长度
	// start：选择nums下标的起点
	// c：当前已选取的元素集合
	if len(c) == k { // 当前选择的元素集合长度 已经等同 设定的子集长度
		*ans = append(*ans, append([]int(nil), c...))
		return
	}
	// 剪枝
	for i := start; i < len(nums)-(k-len(c))+1; i++ { // i的大小不会超过nums长度
		c = append(c, nums[i]) // 选择当前元素
		generateSubsets(nums, k, i+1, c, ans)
		c = c[:len(c)-1] // 回溯，不选当前元素，没有进行递归调用
	}
	return
}

// 回溯
func subsets1(nums []int) (ans [][]int) {
	set := []int{}

	var dfs func(int)
	dfs = func(cur int) { // cur：选到nums中下标为cur的元素
		if cur == len(nums) { // 已经选到nums中最后一个元素
			ans = append(ans, append([]int(nil), set...))
			return
		}

		// 选
		set = append(set, nums[cur])
		dfs(cur + 1)
		// 不选，回溯
		set = set[:len(set)-1]
		dfs(cur + 1)
	}

	dfs(0)
	return
}
