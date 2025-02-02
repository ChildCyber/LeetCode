package leetcode

// 划分字母区间
// https://leetcode.cn/problems/partition-labels/

// 贪心
// 一个片段的结束点必须至少等于该片段内所有字母的最后出现位置的最大值。只有这样，才能确保所有字母都包含在一个片段内。
func partitionLabels(S string) []int {
	var lastIndexOf [26]int
	for i, v := range S {
		lastIndexOf[v-'a'] = i
	}
	var arr []int
	for start, end := 0, 0; start < len(S); start = end + 1 {
		end = lastIndexOf[S[start]-'a']
		for i := start; i < end; i++ {
			if end < lastIndexOf[S[i]-'a'] {
				end = lastIndexOf[S[i]-'a']
			}
		}
		arr = append(arr, end-start+1)
	}
	return arr
}

func partitionLabels1(s string) []int {
	// 记录每个字符的最后出现位置
	lastOccurrence := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		lastOccurrence[s[i]] = i
	}

	var result []int
	start, end := 0, 0
	// 跟踪每个字母的最后出现位置，确保段内所有字母都不会出现在后续段中
	for i := 0; i < len(s); i++ {
		// 更新当前片段的结束位置
		if lastOccurrence[s[i]] > end {
			end = lastOccurrence[s[i]]
		}
		// 贪心地扩展当前段的边界，直到不能再扩展（即到达当前段的结束点），保证段的最小长度，从而最大化段的数量
		// 当当前索引等于结束位置时，分割片段
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}
	return result
}
