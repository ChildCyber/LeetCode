package leetcode

import "strconv"

// 有效的数独
// https://leetcode.cn/problems/valid-sudoku
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

// 一次遍历
// 使用哈希表进行计数，记录每一行、每一列和每一个小九宫格中，每个数字出现的次数
func isValidSudoku(board [][]byte) bool {
	// 创建二维数组 rows 和 columns 分别记录数独的每一行和每一列中的每个数字的出现次数
	var rows, columns [9][9]int // 0～n行，每行n个元素，元素表示第n行的当前数字是否出现过，1为出现，0为未出现
	// 创建三维数组 subboxes 记录数独的每一个小九宫格中的每个数字的出现次数
	var subboxes [3][3][9]int
	// rows[i][index]、columns[j][index] 和 subboxes[i/3][j/3][index] 分别表示数独的第 i 行第 j 列的单元格所在的行、列和小九宫格中
	// 数字 index+1 出现的次数，其中 0≤index<9，对应的数字 index+1 满足 1≤index+1≤9

	// 遍历数独，遍历的过程中更新哈希表中的计数，并判断是否满足有效的数独的条件
	for i, row := range board { // 行
		for j, c := range row { // 列
			if c == '.' {
				continue
			}

			index := c - '1' // board[i][j]对应的数字
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++

			// 判断是否行、列、3×3中首次出现
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}

	return true
}
