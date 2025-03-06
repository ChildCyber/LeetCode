package leetcode

import "strings"

// 简化路径
// https://leetcode.cn/problems/simplify-path/

// 栈
func simplifyPath(path string) string {
	// 1. 按 '/' 切割
	parts := strings.Split(path, "/")
	stack := make([]string, 0)

	// 2. 从前向后遍历
	for _, part := range parts {
		if part == "" || part == "." {
			// 跳过空字符串和当前目录
			continue
		}
		if part == ".." {
			// 上级目录：弹出栈顶元素（如果栈非空）
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			// 普通目录名：压入栈中
			stack = append(stack, part)
		}
	}

	// 3. 构建结果
	return "/" + strings.Join(stack, "/")
}
