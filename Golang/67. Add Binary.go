package leetcode

// 二进制求和
// https://leetcode-cn.com/problems/add-binary/

// 模拟
func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	ans := []byte{}

	// 从右向左遍历，直到处理完所有位和进位
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		// 加上a的当前位
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		// 加上b的当前位
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// 计算当前位和进位
		ans = append(ans, byte(sum%2)+'0')
		carry = sum / 2
	}

	// 反转结果（因为是从低位到高位计算的）
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}

	return string(ans)
}
