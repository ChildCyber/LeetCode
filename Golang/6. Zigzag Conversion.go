package leetcode

import "strings"

// Z 字形变换
// https://leetcode.cn/problems/zigzag-conversion/

// 直接构造法
func convertMatrix(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	n := len(s)
	grid := make([][]byte, numRows)
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := range grid[i] {
			grid[i][j] = ' ' // 填空格方便区分
		}
	}

	r, c := 0, 0 // r表示当前行，c表示当前列
	down := true // 表示方向（往下 or 往右上）

	for i := 0; i < n; i++ {
		grid[r][c] = s[i]
		if down {
			if r == numRows-1 { // 在最底行 → 改为往右上
				down = false
				r--
				c++
			} else { // 向下
				r++
			}
		} else {
			if r == 0 { // 在最顶行 → 改为往下
				down = true
				r++
			} else { // 向右上
				r--
				c++
			}
		}
	}

	var ans strings.Builder
	for i := 0; i < numRows; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != ' ' {
				ans.WriteByte(grid[i][j])
			}
		}
	}
	return ans.String()
}

// 行号分配
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)
	curRow := 0        // 当前行号
	goingDown := false // 是否正在往下走

	for _, ch := range s {
		// 每次把当前字符加进rows[curRow]
		rows[curRow] += string(ch)
		// 如果走到最上面（0）或最下面（numRows-1），就反转方向
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		// 根据方向调整curRow
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	var ans strings.Builder
	for _, row := range rows {
		ans.WriteString(row)
	}

	return ans.String()
}
