package leetcode

// 不同的二叉搜索树
// https://leetcode-cn.com/problems/unique-binary-search-trees/

// 动态规划
// https://leetcode.cn/problems/unique-binary-search-trees/solution/buton-by-ao-zi-ge-pilg/
func numTreesDP(n int) int {
	// 给定一个有序序列1⋯n，可以遍历每个数字 i，将该数字作为树根，将1⋯(i−1)序列作为左子树，将(i+1)⋯n序列作为右子树
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1 // 空树也是树
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// 数学
// 卡特兰数
func numTrees(n int) int {
	ans := 1
	for i := 0; i < n; i++ {
		ans = ans * 2 * (2*i + 1) / (i + 2)
	}
	return ans
}
