package leetcode

// 除自身以外数组的乘积
// https://leetcode-cn.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	length := len(nums)
	// L 和 R 分别表示左右两侧的乘积列表
	l, r, answer := make([]int, length), make([]int, length), make([]int, length)

	// L[i] 为索引 i 左侧所有元素的乘积
	// 对于索引为 '0' 的元素，因为左侧没有元素，所以 L[0] = 1
	l[0] = 1
	for i := 1; i < length; i++ {
		l[i] = nums[i-1] * l[i-1]
	}
	// R[i] 为索引 i 右侧所有元素的乘积
	// 对于索引为 'length-1' 的元素，因为右侧没有元素，所以 R[length-1] = 1
	r[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}
	// 对于索引 i，除 nums[i] 之外其余各元素的乘积就是左侧所有元素的乘积乘以右侧所有元素的乘积
	for i := 0; i < length; i++ {
		answer[i] = l[i] * r[i]
	}
	return answer
}