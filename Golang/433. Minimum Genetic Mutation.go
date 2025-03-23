package leetcode

// 最小基因变化
// https://leetcode.cn/problems/minimum-genetic-mutation/

// 问题本质：隐式图上的单源最短路径问题
// 图的抽象：
//   顶点：每一个有效的基因字符串（startGene和bank中的所有基因）都是一个节点
//   边：如果两个基因字符串恰好只有一个字符不同（即只需要进行一次变异），那么它们之间就有一条边
// 目标：
//   寻找从startGene节点到endGene节点的最小变异次数 → 最短路径

// BFS
func minMutation(startGene string, endGene string, bank []string) int {
	// 将bank转换为集合，便于快速查找
	bankSet := make(map[string]bool)
	for _, gene := range bank {
		bankSet[gene] = true
	}

	// 如果目标基因不在基因库中，直接返回-1
	if !bankSet[endGene] {
		return -1
	}

	// 如果起始基因就是目标基因
	if startGene == endGene {
		return 0
	}

	// 可能的字符
	chars := []byte{'A', 'C', 'G', 'T'}

	// BFS初始化
	queue := []string{startGene}
	visited := make(map[string]bool)
	visited[startGene] = true
	mutations := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			// 如果到达目标基因
			if current == endGene {
				return mutations
			}

			// 尝试改变每一个位置
			geneBytes := []byte(current)
			for j := 0; j < len(geneBytes); j++ {
				originalChar := geneBytes[j]

				// 尝试替换为其他三种字符
				for _, char := range chars {
					if char == originalChar {
						continue
					}

					geneBytes[j] = char
					newGene := string(geneBytes)

					// 检查新基因是否在基因库中且未被访问
					if bankSet[newGene] && !visited[newGene] {
						visited[newGene] = true
						queue = append(queue, newGene)
					}
				}

				// 恢复原字符
				geneBytes[j] = originalChar
			}
		}
		mutations++
	}

	return -1
}
