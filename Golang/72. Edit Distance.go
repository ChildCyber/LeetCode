package leetcode

// 编辑距离
// https://leetcode.cn/problems/edit-distance/

// 暴力
// 递归尝试所有可能的操作序列，找出操作次数最少的那一个
func minDistanceBrute(word1, word2 string) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 基准情况：任一字符串遍历完成
		if i == len(word1) {
			return len(word2) - j // 需要插入剩余字符
		}
		if j == len(word2) {
			return len(word1) - i // 需要删除剩余字符
		}

		// 当前字符匹配：直接比较下一个字符
		if word1[i] == word2[j] {
			return dfs(i+1, j+1)
		}

		// 当前字符不匹配：三种操作选择最小值
		// 1. 插入：在word1的i位置插入word2[j]，然后比较word1[i]和word2[j+1]
		insertOp := dfs(i, j+1)
		// 2. 删除：删除word1[i]，然后比较word1[i+1]和word2[j]
		deleteOp := dfs(i+1, j)
		// 3. 替换：将word1[i]替换为word2[j]，然后比较word1[i+1]和word2[j+1]
		replaceOp := dfs(i+1, j+1)
		return 1 + minThree(insertOp, deleteOp, replaceOp)
	}
	return dfs(0, 0)
}

// 动态规划
// 时间复杂度：O(m×n)
// 空间复杂度：O(m×n)
func minDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	// 有一个字符串为空串
	if m*n == 0 {
		return m + n
	}

	// 状态定义：dp[i][j] 表示将 word1 的前 i 个字符转换为 word2 的前 j 个字符所需的最少操作次数
	// 状态转移方程：
	//  情况1：字符匹配（无需操作）
	//    dp[i][j] = dp[i-1][j-1]
	//    如果当前字符相同，那么编辑距离等于去掉这两个字符后的子问题的编辑距离。
	//  情况2：字符不匹配（需要操作）
	//    dp[i][j] = 1 + min(
	//    	dp[i][j-1],    // 插入操作：在word1中插入word2[j]，相当于word1的前i个字符匹配word2的前j-1个字符
	//    	dp[i-1][j],    // 删除操作：删除word1[i]，相当于word1的前i-1个字符匹配word2的前j个字符
	//    	dp[i-1][j-1]   // 替换操作：将word1[i]替换为word2[j]，相当于word1的前i-1个字符匹配word2的前j-1个字符
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// 边界状态初始化
	for i := 0; i <= m; i++ {
		dp[i][0] = i // 删除所有字符
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // 插入所有字符
	}

	// 计算所有DP值
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				// 字符匹配，无需操作
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 字符不匹配，取三种操作的最小值 + 1
				dp[i][j] = 1 + minThree(
					dp[i][j-1],   // 插入
					dp[i-1][j],   // 删除
					dp[i-1][j-1], // 替换
				)
			}
		}
	}
	return dp[m][n]
}

// 三数取最小值
func minThree(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
