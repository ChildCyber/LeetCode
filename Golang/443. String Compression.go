package leetcode

import "strconv"

// 压缩字符串
// https://leetcode.cn/problems/string-compression/

// 快慢指针
func compress(chars []byte) int {
	n := len(chars)
	if n == 0 {
		return 0
	}

	writeIndex := 0 // 慢指针：负责在数组前面写入压缩后的结果
	readIndex := 0  // 快指针：负责遍历整个字符数组，统计连续相同字符的数量

	for readIndex < n {
		currentChar := chars[readIndex]
		count := 0

		// 统计连续相同字符的数量
		for readIndex < n && chars[readIndex] == currentChar {
			readIndex++
			count++
		}

		// 写入当前字符
		chars[writeIndex] = currentChar
		writeIndex++

		// 如果计数大于1，写入计数
		if count > 1 {
			// 将计数转换为字符并写入
			countStr := strconv.Itoa(count)
			for i := 0; i < len(countStr); i++ {
				chars[writeIndex] = countStr[i]
				writeIndex++
			}
		}
	}

	return writeIndex
}
