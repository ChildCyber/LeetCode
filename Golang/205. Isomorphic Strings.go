package leetcode

// 同构字符串
// https://leetcode.cn/problems/isomorphic-strings/

// 哈希表-双向映射
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 创建两个映射表
	sToT := make(map[byte]byte) // s->t的映射
	tToS := make(map[byte]byte) // t->s的映射

	for i := 0; i < len(s); i++ {
		charS, charT := s[i], t[i]

		// 检查s->t映射是否一致
		if mappedT, ok := sToT[charS]; ok {
			if mappedT != charT {
				return false
			}
		} else {
			// 建立新的s->t映射
			sToT[charS] = charT
		}

		// 检查t->s映射是否一致
		if mappedS, ok := tToS[charT]; ok {
			if mappedS != charS {
				return false
			}
		} else {
			// 建立新的t->s映射
			tToS[charT] = charS
		}
	}

	return true
}
