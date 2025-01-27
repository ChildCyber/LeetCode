package leetcode

// 轮转数组
// https://leetcode.cn/problems/rotate-array/
// 使用额外的数组-切片操作
func rotate1(nums []int, k int) {
	newNums := make([]int, 0)
	k %= len(nums)
	// 数组分成两部分：从开始到倒数第k个元素和最后k个元素
	// 将最后k个元素放在前面，剩下的部分放在后面，形成新的数组
	newNums = append(newNums, nums[len(nums)-k:]...)
	newNums = append(newNums, nums[:len(nums)-k]...)
	copy(nums, newNums)
}

// 使用额外的数组-复制
func rotate2(nums []int, k int) {
	n := len(nums)
	newNums := make([]int, n)
	// 遍历原数组，将原数组下标为 i 的元素放至新数组下标为 (i+k) % n 的位置
	for i, v := range nums {
		newNums[(i+k)%n] = v
	}
	copy(nums, newNums)
}

// 数组翻转-原地三次反转
func rotate(nums []int, k int) {
	k %= len(nums)
	// 所有元素翻转，后 k 个元素移到了前面，但顺序颠倒
	reverse189(nums)
	// 翻转前 k 个元素，纠正了前部分的顺序
	reverse189(nums[:k])
	// 翻转剩余元素，纠正了后部分的顺序
	reverse189(nums[k:])
}

func reverse189(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 环状替换
