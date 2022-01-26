package leetcode

import "sort"

// 分发饼干
// https://leetcode-cn.com/problems/assign-cookies/
// 排序+贪心
/*func findContentChildren(greed, size []int) (ans int) {
	sort.Ints(greed)
	sort.Ints(size)
	n, m := len(greed), len(size)
	for i, j := 0, 0; i < n && j < m; i++ { // 胃口
		for j < m && greed[i] > size[j] {
			j++
		}
		if j < m {
			ans++
			j++
		}
	}
	return
}*/

func findContentChildren1(greed []int, size []int) int {
	sort.Ints(greed)
	sort.Ints(size)
	gi, si, res := 0, 0, 0
	for gi < len(greed) && si < len(size) {
		if size[si] >= greed[gi] { // 满足胃口
			res++
			si++
			gi++
		} else {
			si++
		}
	}
	return res
}
