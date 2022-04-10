package leetcode

// 下一个排列
// https://leetcode-cn.com/problems/next-permutation/
// 两遍扫描
// 有 3 个问题需要解决：
// 如何找到下一个排列？不存在下一个排列的时候如何生成最小的排列？如何原地修改？
// 将一个左边的「较小数」与一个右边的「较大数」交换，以能够让当前排列变大，从而得到下一个排列
// 要让这个「较小数」尽量靠右，而「较大数」尽可能小
// 当交换完成后，「较大数」右边的数需要按照升序重新排列。这样可以在保证新排列大于原来排列的情况下，使变大的幅度尽可能小
func nextPermutation(nums []int) {
	// 对于长度为 n 的排列 nums：
	// 首先从后向前查找第一个顺序对 (i,i+1)，满足 a[i]<a[i+1]。这样「较小数」即为 nums[i]。此时 [i+1,n) 必然是下降序列。
	// 如果找到了顺序对，那么在区间 [i+1,n) 中从后向前查找第一个元素 j 满足 nums[i]<nums[j]。这样「较大数」即为 nums[j]。
	// 交换 nums[i] 与 nums[j]，此时可以证明区间 [i+1,n) 必为降序。可以直接使用双指针反转区间 [i+1,n) 使其变为升序，而无需对该区间进行排序。
	n := len(nums)
	i := n - 2                           // 从末尾向前扫描，下标从n-1开始，最后一个元素跳过，所以从n-2开始
	for i >= 0 && nums[i] >= nums[i+1] { // 寻找非降序的nums[i-1]，从后向前，nums中[i,n)部分为降序
		i--
	}
	// 寻找nums[j]
	if i >= 0 { // 判断i是否越界（重点）
		j := n - 1                         // 从末尾向前扫描，下标从n-1开始
		for j >= 0 && nums[i] >= nums[j] { // nums[j]要大于nums[i]（重点）
			j--
		}
		nums[i], nums[j] = nums[j], nums[i] // 交换i，j位置上的数字（重点）
	}
	// i有可能为-1，即不存在下一个更大的排列，那么这个数组会重排为字典序最小的排列，反转整个数组
	reverse(nums[i+1:]) // [i+1,n)为降序，直接反转为升序
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
