package leetcode

import (
	"math/rand"
	"sort"
)

// 按权重随机选择
// https://leetcode-cn.com/problems/random-pick-with-weight/
// 前缀和 + 二分查找
type Solution struct {
	pre []int
}

func Constructor528(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (s *Solution) PickIndex() int {
	x := rand.Intn(s.pre[len(s.pre)-1]) + 1
	return sort.SearchInts(s.pre, x)
}
