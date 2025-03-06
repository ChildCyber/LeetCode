package leetcode

// 赎金信
// https://leetcode.com/problems/ransom-note/

// 哈希表
func canConstruct(ransomNote string, magazine string) bool {
	// 创建字符计数的哈希表
	charCount := make(map[byte]int)

	// 统计magazine中每个字符的出现次数
	for i := 0; i < len(magazine); i++ {
		charCount[magazine[i]]++
	}

	// 检查ransomNote中的每个字符
	for i := 0; i < len(ransomNote); i++ {
		ch := ransomNote[i]
		// 如果字符不存在或计数已为0，返回false
		if charCount[ch] <= 0 {
			return false
		}
		// 减少对应字符的计数
		charCount[ch]--
	}

	return true
}

// 数组
func canConstructArray(ransomNote string, magazine string) bool {
	// 使用数组代替map
	charCount := [26]int{}

	// 统计magazine中字符出现次数
	for _, ch := range magazine {
		charCount[ch-'a']++
	}

	// 检查ransomNote中的字符
	for _, ch := range ransomNote {
		index := ch - 'a'
		if charCount[index] <= 0 {
			return false
		}
		charCount[index]--
	}

	return true
}
