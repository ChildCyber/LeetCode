package leetcode

// 全排列
// https://leetcode-cn.com/problems/permutations/
// 回溯， DFS深搜，树形结构
// 时间复杂度：O(n×n!)，其中 n 为序列的长度
// 空间复杂度：O(n)，其中 n 为序列的长度
func permute(nums []int) [][]int { // 不含重复数字的数组
	if len(nums) == 0 {
		return [][]int{}
	}

	// 状态变量
	// path:已经选了哪些数，数组
	// depth:递归到第几层
	// used布尔数组
	// res:结果
	used, path, res := make([]bool, len(nums)), []int{}, [][]int{}
	generatePermutation(nums, 0, path, &res, &used)
	return res
}

func generatePermutation(nums []int, depth int /*递归深度*/, path []int, res *[][]int, used *[]bool) {
	// 当前递归层数等于数组长度，递归终止
	if depth == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	// 递归，考虑数组中所有的数
	for i := 0; i < len(nums); i++ { // 从下标0开始
		if !(*used)[i] { // 当前数未被使用，使用该数继续进行回溯
			(*used)[i] = true                                   // 标记为已使用
			path = append(path, nums[i])                        // 当前数添加到path
			generatePermutation(nums, depth+1, path, res, used) // 递归考虑下一层，只有depth需要变，表示递归到下一层
			// 回溯部分，前面做了什么，就需要做对应的逆操作
			path = path[:len(path)-1] // 回到上一层节点
			(*used)[i] = false        // 未使用
		}
	}
	return
}
