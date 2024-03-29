package leetcode

import (
	"strconv"
)

// hard
// 排列序列
// https://leetcode-cn.com/problems/permutation-sequence/
// 回溯
func getPermutation(n int, k int) string {
	if k == 0 {
		return ""
	}
	used, p, res := make([]bool, n), []int{}, ""
	findPermutation(n, 0, &k, p, &res, &used)
	return res
}

func findPermutation(n, index int, k *int, p []int, res *string, used *[]bool) {
	if index == n {
		*k--
		if *k == 0 {
			for _, v := range p {
				*res += strconv.Itoa(v + 1)
			}
		}
		return
	}
	for i := 0; i < n; i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, i)
			findPermutation(n, index+1, k, p, res, used)
			// 回溯
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	// 剪枝
	return
}
