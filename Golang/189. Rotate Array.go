package leetcode

// 轮转数组
// https://leetcode.cn/problems/rotate-array/
// 使用额外的数组
func rotate1(nums []int, k int) {
	newNums := make([]int, 0)
	newNums = append(newNums, nums[len(nums)-k:]...)
	newNums = append(newNums, nums[:k+1]...)
	copy(nums, newNums)
}

func rotate2(nums []int, k int) {
	n := len(nums)
	newNums := make([]int, n)
	// 遍历原数组，将原数组下标为 i 的元素放至新数组下标为 (i+k) mod n 的位置
	for i, v := range nums {
		newNums[(i+k)%n] = v
	}
	copy(nums, newNums)
}

// 数组翻转
func reverse189(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	// 所有元素翻转
	reverse189(nums)
	// 翻转[0,k mod n−1]区间的元素
	reverse189(nums[:k])
	// 翻转[k mod n,n−1]区间的元素
	reverse189(nums[k:])
}
