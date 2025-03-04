package leetcode

// 反转字符串中的元音字母
// https://leetcode.cn/problems/reverse-vowels-of-a-string/

// 对撞指针
func reverseVowels(s string) string {
	bytes := []byte(s)
	left, right := 0, len(bytes)-1

	// 定义元音字母集合（包括大小写）
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	for left < right {
		// 移动左指针直到找到元音
		for left < right && !vowels[bytes[left]] {
			left++
		}

		// 移动右指针直到找到元音
		for left < right && !vowels[bytes[right]] {
			right--
		}

		// 交换两个元音
		if left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}

	return string(bytes)
}
