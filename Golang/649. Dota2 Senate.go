package leetcode

// Dota2 参议院
// https://leetcode.cn/problems/dota2-senate/

// 队列
func predictPartyVictory(senate string) string {
	n := len(senate)
	radiant := make([]int, 0)
	dire := make([]int, 0)

	// 初始化队列，存储参议员的索引
	for i, ch := range senate {
		if ch == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	// 模拟投票过程
	for len(radiant) > 0 && len(dire) > 0 {
		rIndex := radiant[0]
		dIndex := dire[0]

		// 移除队首元素
		radiant = radiant[1:]
		dire = dire[1:]

		if rIndex < dIndex {
			// Radiant先投票，禁止Dire，Radiant重新入队
			radiant = append(radiant, rIndex+n)
		} else {
			// Dire先投票，禁止Radiant，Dire重新入队
			dire = append(dire, dIndex+n)
		}
	}

	// 判断获胜方
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
