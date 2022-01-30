package leetcode

import "strconv"

// 复原 IP 地址
// https://leetcode-cn.com/problems/restore-ip-addresses/
// 回溯
func restoreIPAddresses(s string) []string {
	if s == "" {
		return []string{}
	}
	res, ip := []string{}, []int{}
	dfs93(s, 0, ip, &res)
	return res
}

func dfs93(s string, index int, ip []int, res *[]string) {
	// 找到了 4 段 IP 地址并且遍历完了字符串
	if index == len(s) {
		if len(ip) == 4 {
			*res = append(*res, getString93(ip))
		}
		return
	}
	if index == 0 {
		num, _ := strconv.Atoi(string(s[0]))
		ip = append(ip, num)
		dfs93(s, index+1, ip, res)
	} else {
		num, _ := strconv.Atoi(string(s[index]))
		next := ip[len(ip)-1]*10 + num
		if next <= 255 && ip[len(ip)-1] != 0 {
			ip[len(ip)-1] = next
			dfs93(s, index+1, ip, res)
			ip[len(ip)-1] /= 10
		}
		if len(ip) < 4 {
			ip = append(ip, num)
			dfs93(s, index+1, ip, res)
			ip = ip[:len(ip)-1]
		}
	}
}

func getString93(ip []int) string {
	res := strconv.Itoa(ip[0])
	for i := 1; i < len(ip); i++ {
		res += "." + strconv.Itoa(ip[i])
	}
	return res
}
