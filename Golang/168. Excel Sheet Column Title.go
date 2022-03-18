package leetcode

// Excel表列名称
// https://leetcode-cn.com/problems/excel-sheet-column-title/
// 数学
/*从 1 开始的的 26 进制转换题
对于一般性的进制转换题目，只需要不断地对 columnNumber 进行 % 运算取得最后一位，然后对 columnNumber 进行 / 运算，将已经取得的位数去掉，直到 columnNumber 为 0 即可。
从 1 开始，因此在执行「进制转换」操作前，需要先对 columnNumber 执行减一操作，从而实现整体偏移。
*/
func convertToTitle(n int) string {
	ans := []byte{}
	for n > 0 {
		ans = append(ans, 'A'+byte((n-1)%26))
		n = (n - 1) / 26 // 下一位，26进制
	}
	// 反转int数组
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
