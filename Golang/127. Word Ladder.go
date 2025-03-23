package leetcode

// 单词接龙
// https://leetcode.cn/problems/word-ladder/description/

// 问题本质：隐式图上的单源最短路径问题
// 图的抽象：
//   顶点：所有的单词（包括 beginWord 和 wordList 中的所有单词）
//   边：如果两个单词只有一个字母不同，则存在一条无向边
// 目标：
//   找到从beginWord到endWord的最短转换序列的长度 → 最短路径

// BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	level := 1

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			if current == endWord {
				return level
			}

			// 生成所有可能的邻居
			chars := []byte(current)
			for j := 0; j < len(chars); j++ {
				originalChar := chars[j]

				// 尝试改变每一个位置
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == originalChar {
						continue
					}

					chars[j] = byte(c)
					newWord := string(chars)

					if wordSet[newWord] && !visited[newWord] {
						visited[newWord] = true
						queue = append(queue, newWord)
					}
				}

				chars[j] = originalChar
			}
		}
		level++
	}

	return 0
}
