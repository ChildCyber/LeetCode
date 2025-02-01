package leetcode

// 下一个排列
// https://leetcode-cn.com/problems/next-permutation/
// 两遍扫描
func nextPermutation(nums []int) {
	// 对于长度为 n 的排列 nums：
	// 1. 寻找非降序的nums[i-1]
	// 从后向前查找第一个顺序对 (i,i+1)，满足 a[i]<a[i+1]。这样「较小数」即为 nums[i]。此时 [i+1,n) 必然是下降序列。
	// 2. 寻找nums[j]
	// 如果找到了顺序对，那么在区间 [i+1,n) 中从后向前查找第一个元素 j 满足 nums[i]<nums[j]。这样「较大数」即为 nums[j]。
	// 3. 交换nums[i-1]和nums[j]
	// 交换 nums[i] 与 nums[j]，此时可以证明区间 [i+1,n) 必为降序。可以直接使用双指针反转区间 [i+1,n) 使其变为升序，而无需对该区间进行排序。
	// 4. 反转nums[i-1]后的元素
	n := len(nums)

	// 步骤1：从右向左找到第一个下降点
	i := n - 2                           // 从末尾向前扫描，下标从n-1开始，最后一个元素跳过，所以从n-2开始
	for i >= 0 && nums[i] >= nums[i+1] { // 寻找非降序的nums[i-1]，从后向前，nums中[i,n)部分为降序
		i--
	}

	// 步骤2：如果找到下降点，找到比下降点数字大的最小数字
	if i >= 0 { // 判断i是否越界（重点）
		// 从右向左找到第一个大于nums[i]的数
		j := n - 1
		for j > i && nums[i] >= nums[j] { // nums[j]要大于nums[i]（重点）
			j--
		}
		// 交换i，j位置上的数字（重点）
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 步骤3：反转i+1到末尾的部分。[i+1,n)为降序，直接反转为升序
	reverse(nums[i+1:]) // i有可能为-1，即不存在下一个更大的排列，那么这个数组会重排为字典序最小的排列，反转整个数组
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
