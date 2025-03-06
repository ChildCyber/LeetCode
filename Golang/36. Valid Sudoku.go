package leetcode

import "strconv"

// 有效的数独
// https://leetcode.cn/problems/valid-sudoku/

// 暴力
// 时间复杂度：O(n^3)
func isValidSudokuForce(board [][]byte) bool {
	// 判断行 row，每行只出现一次
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := board[i][j]
			if string(cellVal) != "." {
				index, _ := strconv.Atoi(string(cellVal))
				if index > 9 || index < 1 { // 无效数字
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}

	// 判断列 column，每列只出现一次
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := board[j][i]
			if string(cellVal) != "." {
				index, _ := strconv.Atoi(string(cellVal))
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}

	// 判断 9宫格 3X3 cell，3×3只出现一次
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := [10]int{}
			for ii := i * 3; ii < i*3+3; ii++ {
				for jj := j * 3; jj < j*3+3; jj++ {
					cellVal := board[ii][jj]
					if string(cellVal) != "." {
						index, _ := strconv.Atoi(string(cellVal))
						if tmp[index] == 1 {
							return false
						}
						tmp[index] = 1
					}
				}
			}
		}
	}

	return true
}

// 哈希表
// 一次遍历：使用哈希表进行计数，记录每一行、每一列和每一个小九宫格中，每个数字出现的次数
func isValidSudoku(board [][]byte) bool {
	// rows[9][9]：第i行中数字j是否出现过
	// cols[9][9]：第i列中数字j是否出现过
	// boxes[9][9]：第k个3×3宫格中数字j是否出现过
	var rows, cols, boxes [9][9]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			ch := board[r][c]
			if ch == '.' {
				continue
			}
			num := ch - '1'               // 数字索引 0~8
			boxIndex := (r/3)*3 + (c / 3) // 定位宫格编号

			if rows[r][num] || cols[c][num] || boxes[boxIndex][num] {
				return false
			}

			rows[r][num] = true
			cols[c][num] = true
			boxes[boxIndex][num] = true
		}
	}
	return true
}
