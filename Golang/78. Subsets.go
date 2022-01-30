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
		// 回溯
		c = c[:len(c)-1]
	}
	return
}
