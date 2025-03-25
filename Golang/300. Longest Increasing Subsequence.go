package leetcode

// 最长递增子序列
// https://leetcode-cn.com/problems/longest-increasing-subsequence/

// 暴力
// 通过回溯获取数组的所有子序列，再对这些子串再依次判定是否为「严格上升」
func lengthOfLISBacktrace(nums []int) int {
	// dfs437 O(N*2^n)
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

// 动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 状态定义：dp[i] 表示以 nums[i] 这个元素结尾的最长递增子序列的长度（包含第 i 个元素作为结尾的最长子序列长度）
	// 状态转移方程：dp[i] = max(dp[i], dp[j]+1)，对于所有 j < i && nums[j] < nums[i]
	dp, ans := make([]int, len(nums)), 1

	// 初始化dp数组，每个位置至少长度为1（自身）
	for i := range dp {
		dp[i] = 1
	}

	// 动态规划计算每个位置的最长递增子序列长度
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// 如果nums[j] < nums[i]，说明可以将nums[i]接在nums[j]后面
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i]) // 记录最长上升子序列的长度
	}
	return ans
}

func lengthOfLISDP(nums []int) int {
	ans, length := 0, len(nums)
	if length == 0 {
		return ans
	}
	dp := make([]int, length) // dp长度和原数组相同
	for i := 0; i < length; i++ {
		dp[i] = 1                // 每一个数都有可能是最长子序列的起点
		for j := 0; j < i; j++ { // j < i
			if nums[i] > nums[j] { // nums[0:j]逐个和nums[i]进行比较，获取上升子序列
				dp[i] = max(dp[i], dp[j]+1)
			}
			ans = max(dp[i], ans)
		}
	}
	return ans
}

// 贪心+二分查找
// 时间复杂度：O(n log n)
// 空间复杂度：O(n)
func lengthOfLISBS(nums []int) int {
	// 维护一个数组 tails，其中 tails[i] 表示长度为 i+1 的所有递增子序列中，最小的末尾元素值
	if len(nums) == 0 {
		return 0
	}

	tails := make([]int, 0)

	// 遍历原数组中的每个元素：
	//   如果当前元素大于 tails 的最后一个元素，直接添加到末尾
	//   否则，在 tails 中找到第一个大于等于当前元素的位置，用当前元素替换它
	for _, num := range nums {
		// 二分查找插入位置
		left, right := 0, len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 如果找到的位置等于tails长度，说明num比所有元素都大
		if left == len(tails) {
			tails = append(tails, num)
		} else {
			// 否则替换该位置的元素
			tails[left] = num
		}
	}

	return len(tails) // tails的长度就是最长递增子序列的长度
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
