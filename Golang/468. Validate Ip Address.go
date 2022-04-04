package leetcode

import (
	"strconv"
	"strings"
)

// 验证IP地址
// https://leetcode.cn/problems/validate-ip-address/
// 依次判断
func validIPAddress(queryIP string) string {
	// IPv4
	if sp := strings.Split(queryIP, "."); len(sp) == 4 {
		for _, s := range sp {
			if len(s) > 1 && s[0] == '0' { // 是否不包含前导零
				return "Neither"
			}
			if v, err := strconv.Atoi(s); err != nil || v > 255 { // 是否只包含数字，值是否在 [0,255] 之间
				return "Neither"
			}
		}
		return "IPv4"
	}
	// IPv6
	if sp := strings.Split(queryIP, ":"); len(sp) == 8 {
		for _, s := range sp {
			if len(s) > 4 { // 长度是否在 [1,4] 之间
				return "Neither"
			}
			if _, err := strconv.ParseUint(s, 16, 64); err != nil { // 是否只包含数字，或者 a-f，或者 A-F
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}
