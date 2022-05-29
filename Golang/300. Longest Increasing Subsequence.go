package leetcode

// 最长递增子序列
// https://leetcode-cn.com/problems/longest-increasing-subsequence/
// 动态规划
// O(n^2)
func lengthOfLIS(nums []int) int {
	// dp[i] 代表 [0,i] 范围内,选择数字 nums[i] 为结尾可以获得的最⻓上升子序列的⻓度。
	// 状态转移方程为 dp[i] = max(dp[j]) + 1 ,其中 j < i && nums[j] > nums[i] ,取所有满足条件的最大值。
	dp, ans := make([]int, len(nums)+1), 0 // 多了个dp[0]，长度需要加1
	dp[0] = 0                              // 不选数字，最⻓上升子序列的⻓度为0
	for i := 1; i <= len(nums); i++ {      // 需要两个for
		for j := 1; j < i; j++ { // j<i
			// nums[i]严格大于在它位置之前的某个数，那么 nums[i]就可以接在这个数后面形成一个更长的上升子序列
			// 找出i之前最大的dp[j]
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1     // dp[1]=1
		ans = max(ans, dp[i]) // 记录最长上升子序列的长度
	}
	return ans
}

func lengthOfLIS1(nums []int) int {
	ans, length := 0, len(nums)
	if length == 0 {
		return ans
	}
	dp := make([]int, length) // dp长度和原数组相同
	for i := 0; i < length; i++ {
		dp[i] = 1                // 每一个数都有可能是最长子序列的起点
		for j := 0; j < i; j++ { // j < i
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			ans = max(dp[i], ans)
		}
	}
	return ans
}

// 分治 + 递归
// https://leetcode.cn/problems/longest-increasing-subsequence/solution/300zui-chang-di-zeng-zi-xu-lie-cong-bao-bwso6/
func lengthOfLISForce(nums []int) int {
	ans := 0

	var dfs func(int, int) int
	dfs = func(index int, last int) int {
		if index == len(nums) {
			return 1
		}
		length := dfs(index+1, last)
		if nums[index] > last {
			length = max(length, dfs(index+1, nums[index])+1)
		}
		return length
	}

	for i := 0; i < len(nums); i++ {
		ans = max(ans, dfs(i, nums[i]))
	}
	return ans
}

// 暴力
// 通过回溯获取数组的所有子序列，再对这些子串再依次判定是否为「严格上升」
func lengthOfLISBacktrace(nums []int) int {
	// dfs O(N*2^n)
	stack := make([]int, 0, 16)
	ans := 0

	var dfs func(nums []int, index int)
	dfs = func(nums []int, index int) {
		for i := index; i < len(nums); i++ {
			if nums[i] <= stack[len(stack)-1] {
				continue
			}

			// 要
			stack = append(stack, nums[i])
			dfs(nums, i+1)
			//不要
			stack = stack[:len(stack)-1]
		}
		ans = max(ans, len(stack))
	}

	for i := 0; i < len(nums); i++ {
		stack = append(stack, nums[i])
		dfs(nums, i+1)
		stack = stack[:len(stack)-1]
	}
	return ans
}
