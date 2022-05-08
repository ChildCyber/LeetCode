package leetcode

// 子集
// https://leetcode-cn.com/problems/subsets/
// 回溯
func subsets(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	for k := 0; k <= len(nums); k++ {
		generateSubsets(nums, k, 0, c, &res)
	}
	return res
}

func generateSubsets(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// 剪枝
	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		c = append(c, nums[i])
		generateSubsets(nums, k, i+1, c, res)
		c = c[:len(c)-1] // 回溯
	}
	return
}

// 回溯
func subsets1(nums []int) (ans [][]int) {
	set := []int{}

	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		// 选
		set = append(set, nums[cur])
		dfs(cur + 1)
		// 不选
		set = set[:len(set)-1]
		dfs(cur + 1)
	}

	dfs(0)
	return
}
