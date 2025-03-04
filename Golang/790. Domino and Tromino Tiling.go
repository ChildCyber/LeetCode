package leetcode

// 多米诺和托米诺平铺
// https://leetcode.cn/problems/domino-and-tromino-tiling/

// 动态规划
func numTilings(n int) int {
	const mod int = 1e9 + 7
	// 状态定义：定义四种状态来表示当前列的填充情况：
	// 设 dp[i][s] 表示铺满前 i 列，且第 i+1 列的填充状态为 s 的方案数。
	// 状态 s 有四种情况（用 0-3 表示）：
	//    状态 0：前 i 列已完全铺满，第 i+1 列也是空的
	//    状态 1：前 i 列已完全铺满，第 i+1 列的上格被铺
	//    状态 2：前 i 列已完全铺满，第 i+1 列的下格被铺
	//    状态 3：前 i 列已完全铺满，第 i+1 列完全被铺
	// 状态转移：
	dp := make([][4]int, n+1)
	dp[0][3] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][3]
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % mod
		dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % mod
		dp[i][3] = (((dp[i-1][0]+dp[i-1][1])%mod+dp[i-1][2])%mod + dp[i-1][3]) % mod
	}
	return dp[n][3]
}

func numTilings1(n int) int {
	const MOD = 1_000_000_007

	// dp[i][mask]: 处理完前 i 列后，列 i+1 的预占为 mask
	dp := make([][4]int64, n+1)
	dp[0][0] = 1 // 初始：列1（即 i=1）上下都空

	// 转移表，按照上面列出的规则
	trans := map[int][]int{
		0: {0, 1, 2, 3},
		1: {2, 3},
		2: {1, 3},
		3: {0},
	}

	for i := 0; i < n; i++ {
		for m := 0; m < 4; m++ {
			v := dp[i][m]
			if v == 0 {
				continue
			}
			for _, nm := range trans[m] {
				dp[i+1][nm] = (dp[i+1][nm] + v) % MOD
			}
		}
	}
	return int(dp[n][0] % MOD)
}
