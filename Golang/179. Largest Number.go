package leetcode

import (
	"sort"
	"strconv"
)

// 最大数
// https://leetcode-cn.com/problems/largest-number/
// 排序
func largestNumber(nums []int) string {
	// 定义比较函数，把最大的放左边
	// 比较两个数不同的拼接顺序的结果，进而决定它们在结果中的排列顺序
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		// 例如[4,42]，比较 442>424
		return sy*x+y > sx*y+x
	})

	if nums[0] == 0 {
		return "0"
	}

	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}

	return string(ans)
}
