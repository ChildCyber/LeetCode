package leetcode

import (
	"strconv"
	"strings"
)

// 相等行列对
// https://leetcode.cn/problems/equal-row-and-column-pairs/

// 暴力
func equalPairsBrute(grid [][]int) int {
	n := len(grid)
	ans := 0

	// 遍历所有行
	for i := 0; i < n; i++ {
		// 遍历所有列
		for j := 0; j < n; j++ {
			match := true
			// 比较第i行和第j列的所有元素
			for k := 0; k < n; k++ {
				if grid[i][k] != grid[k][j] {
					match = false
					break
				}
			}
			if match {
				ans++
			}
		}
	}

	return ans
}

// 哈希表
func equalPairs(grid [][]int) int {
	n := len(grid)
	rowMap := make(map[string]int)

	// 把每一行转成字符串，存频率
	for i := 0; i < n; i++ {
		row := make([]string, n)
		for j := 0; j < n; j++ {
			row[j] = strconv.Itoa(grid[i][j])
		}
		key := strings.Join(row, ",")
		rowMap[key]++
	}

	count := 0
	// 枚举每一列
	for j := 0; j < n; j++ {
		col := make([]string, n)
		for i := 0; i < n; i++ {
			col[i] = strconv.Itoa(grid[i][j])
		}
		key := strings.Join(col, ",")
		if freq, ok := rowMap[key]; ok {
			count += freq
		}
	}

	return count
}
