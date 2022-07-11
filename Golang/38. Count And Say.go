package leetcode

import (
	"strconv"
	"strings"
)

// 外观数列
// https://leetcode.cn/problems/count-and-say
func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ { // 到n
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j { // 遍历字符串prev
			for j < len(prev) && prev[j] == prev[start] { // 出现连续相同的字符，不停向后遍历
				j++
			}
			cur.WriteString(strconv.Itoa(j - start)) // 相同字符连续出现的次数
			cur.WriteByte(prev[start])               // 连续出现的相同字符
		}
		prev = cur.String()
	}
	return prev
}
