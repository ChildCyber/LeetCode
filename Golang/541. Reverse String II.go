package leetcode

// 反转字符串 II
// https://leetcode.cn/problems/reverse-string-ii/

// 模拟
func reverseStr(s string, k int) string {
	bytes := []byte(s)

	for i := 0; i < len(bytes); i += 2 * k {
		// 确定反转区间的右边界
		right := i + k
		if right > len(bytes) {
			right = len(bytes)
		}
		// 反转前k个字符
		reverseBytes(bytes[i:right])
	}

	return string(bytes)
}

// 反转字节切片
func reverseBytes(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}
