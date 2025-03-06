package leetcode

import "strings"

// 翻转字符串里的单词
// https://leetcode-cn.com/problems/reverse-words-in-a-string/

// 标准库
func reverseWords151(s string) string {
	ss := strings.Fields(s)       // 空格拆分字符串
	reverse151(&ss, 0, len(ss)-1) // 翻转
	return strings.Join(ss, " ")  // 连接
}

func reverse151(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}

// 双指针
// 先反转整个字符串，再反转每个单词
func reverseWords(s string) string {
	// 将字符串转换为字节数组以便修改
	bytes := []byte(s)
	n := len(bytes)

	// 辅助函数：反转字节数组中指定范围的字符
	var reverse func([]byte, int, int)
	reverse = func(bytes []byte, left, right int) {
		for left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}

	// 1. 去除多余空格
	// 使用快慢指针去除首尾和中间多余空格
	slow, fast := 0, 0

	// 跳过开头空格
	for fast < n && bytes[fast] == ' ' {
		fast++
	}

	// 处理中间多余空格
	for fast < n {
		if fast > 0 && bytes[fast] == ' ' && bytes[fast-1] == ' ' {
			fast++
			continue
		}
		bytes[slow] = bytes[fast] // fast元素复制到slow
		slow++
		fast++
	}

	// 去除末尾可能的多余空格
	if slow > 0 && bytes[slow-1] == ' ' {
		slow--
	}

	/*for fast < n {
		// 跳过连续的空格
		if bytes[fast] == ' ' {
			fast++
			continue
		}
		// 非空格字符：在单词前添加空格（非第一个单词）
		if slow != 0 {
			bytes[slow] = ' '
			slow++
		}
		// 复制整个单词
		for fast < n && bytes[fast] != ' ' {
			bytes[slow] = bytes[fast]
			slow++
			fast++
		}
	}*/

	// 截取有效部分
	bytes = bytes[:slow]
	n = slow

	// 2. 反转整个字符串
	reverse(bytes, 0, n-1)

	// 3. 反转每个单词
	start := 0
	for i := 0; i <= n; i++ {
		if i == n || bytes[i] == ' ' {
			reverse(bytes, start, i-1)
			start = i + 1
		}
	}

	return string(bytes)
}
