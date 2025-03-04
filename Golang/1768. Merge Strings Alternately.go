package leetcode

// 交替合并字符串
// https://leetcode.cn/problems/merge-strings-alternately/

// 单循环+条件判断
func mergeAlternately(word1 string, word2 string) string {
	n1 := len(word1)
	n2 := len(word2)
	short := min(n1, n2)

	ans := make([]byte, 0, n1+n2)
	for i := 0; i < short; i++ {
		ans = append(ans, word1[i])
		ans = append(ans, word2[i])
	}

	if n1 > n2 {
		ans = append(ans, []byte(word1[n2:])...)
	} else {
		ans = append(ans, []byte(word2[n1:])...)
	}

	return string(ans)
}

// 双指针
func mergeAlternately1(word1 string, word2 string) string {
	n1, n2 := len(word1), len(word2)
	result := make([]byte, n1+n2)

	i, j, k := 0, 0, 0

	// 交替添加字符直到其中一个字符串用完
	for i < n1 && j < n2 {
		result[k] = word1[i]
		k++
		i++
		result[k] = word2[j]
		k++
		j++
	}

	// 添加word1剩余部分
	for i < n1 {
		result[k] = word1[i]
		k++
		i++
	}

	// 添加word2剩余部分
	for j < n2 {
		result[k] = word2[j]
		k++
		j++
	}

	return string(result)
}
